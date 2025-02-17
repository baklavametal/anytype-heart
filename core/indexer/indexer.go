package indexer

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/spacestorage"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"

	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/metrics"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/space/spacecore/storage"
	"github.com/anyproto/anytype-heart/util/slice"
)

const (
	CName = "indexer"
)

var log = logging.Logger("anytype-doc-indexer")

func New() Indexer {
	return &indexer{
		indexedFiles: &sync.Map{},
	}
}

type Indexer interface {
	ForceFTIndex()
	StartFullTextIndex() error
	ReindexMarketplaceSpace(space clientspace.Space) error
	ReindexSpace(space clientspace.Space) error
	RemoveIndexes(spaceId string) (err error)
	Index(ctx context.Context, info smartblock.DocInfo, options ...smartblock.IndexOption) error
	app.ComponentRunnable
}

type Hasher interface {
	Hash() string
}

type indexer struct {
	store          objectstore.ObjectStore
	fileStore      filestore.FileStore
	source         source.Service
	picker         block.ObjectGetter
	ftsearch       ftsearch.FTSearch
	storageService storage.ClientStorage
	fileService    files.Service

	quit       chan struct{}
	btHash     Hasher
	newAccount bool
	forceFt    chan struct{}

	indexedFiles     *sync.Map
	reindexLogFields []zap.Field

	flags reindexFlags
}

func (i *indexer) Init(a *app.App) (err error) {
	i.newAccount = a.MustComponent(config.CName).(*config.Config).NewAccount
	i.store = a.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	i.storageService = a.MustComponent(spacestorage.CName).(storage.ClientStorage)
	i.source = a.MustComponent(source.CName).(source.Service)
	i.btHash = a.MustComponent("builtintemplate").(Hasher)
	i.fileStore = app.MustComponent[filestore.FileStore](a)
	i.ftsearch = app.MustComponent[ftsearch.FTSearch](a)
	i.picker = app.MustComponent[block.ObjectGetter](a)
	i.fileService = app.MustComponent[files.Service](a)
	i.quit = make(chan struct{})
	i.forceFt = make(chan struct{})
	return
}

func (i *indexer) Name() (name string) {
	return CName
}

func (i *indexer) Run(context.Context) (err error) {
	return i.StartFullTextIndex()
}

func (i *indexer) StartFullTextIndex() (err error) {
	if ftErr := i.ftInit(); ftErr != nil {
		log.Errorf("can't init ft: %v", ftErr)
	}
	go i.ftLoop()
	return
}

func (i *indexer) Close(ctx context.Context) (err error) {
	close(i.quit)
	return nil
}

func (i *indexer) Index(ctx context.Context, info smartblock.DocInfo, options ...smartblock.IndexOption) error {
	// options are stored in smartblock pkg because of cyclic dependency :(
	startTime := time.Now()
	opts := &smartblock.IndexOptions{}
	for _, o := range options {
		o(opts)
	}
	err := i.storageService.BindSpaceID(info.Space.Id(), info.Id)
	if err != nil {
		log.Error("failed to bind space id", zap.Error(err), zap.String("id", info.Id))
		return err
	}
	headHashToIndex := headsHash(info.Heads)
	saveIndexedHash := func() {
		if headHashToIndex == "" {
			return
		}

		err = i.store.SaveLastIndexedHeadsHash(info.Id, headHashToIndex)
		if err != nil {
			log.With("objectID", info.Id).Errorf("failed to save indexed heads hash: %v", err)
		}
	}

	indexDetails, indexLinks := info.SmartblockType.Indexable()
	if !indexDetails && !indexLinks {
		return nil
	}

	lastIndexedHash, err := i.store.GetLastIndexedHeadsHash(info.Id)
	if err != nil {
		log.With("object", info.Id).Errorf("failed to get last indexed heads hash: %v", err)
	}

	if opts.SkipIfHeadsNotChanged {
		if headHashToIndex == "" {
			log.With("objectID", info.Id).Errorf("heads hash is empty")
		} else if lastIndexedHash == headHashToIndex {
			log.With("objectID", info.Id).Debugf("heads not changed, skipping indexing")

			// todo: the optimization temporarily disabled to see the metrics
			// return nil
		}
	}

	details := info.Details

	indexSetTime := time.Now()
	var hasError bool
	if indexLinks {
		if err = i.store.UpdateObjectLinks(info.Id, info.Links); err != nil {
			hasError = true
			log.With("objectID", info.Id).Errorf("failed to save object links: %v", err)
		}
	}

	indexLinksTime := time.Now()
	if indexDetails {
		if err := i.store.UpdateObjectDetails(info.Id, details); err != nil {
			if errors.Is(err, objectstore.ErrDetailsNotChanged) {
				metrics.ObjectDetailsHeadsNotChangedCounter.Add(1)
				log.With("objectID", info.Id).With("hashesAreEqual", lastIndexedHash == headHashToIndex).With("lastHashIsEmpty", lastIndexedHash == "").With("skipFlagSet", opts.SkipIfHeadsNotChanged).Debugf("details have not changed")
			} else {
				hasError = true
				log.With("objectID", info.Id).Errorf("can't update object store: %v", err)
			}
		} else {
			// todo: remove temp log
			if lastIndexedHash == headHashToIndex {
				l := log.With("objectID", info.Id).
					With("hashesAreEqual", lastIndexedHash == headHashToIndex).
					With("lastHashIsEmpty", lastIndexedHash == "").
					With("skipFlagSet", opts.SkipIfHeadsNotChanged)

				if opts.SkipIfHeadsNotChanged {
					l.Warnf("details have changed, but heads are equal")
				} else {
					l.Debugf("details have changed, but heads are equal")
				}
			}
		}

		// todo: the optimization temporarily disabled to see the metrics
		if true || !(opts.SkipFullTextIfHeadsNotChanged && lastIndexedHash == headHashToIndex) {
			if err := i.store.AddToIndexQueue(info.Id); err != nil {
				log.With("objectID", info.Id).Errorf("can't add id to index queue: %v", err)
			}
		}

		i.indexLinkedFiles(ctx, info.Space, info.FileHashes)
	} else {
		_ = i.store.DeleteDetails(info.Id)
	}
	indexDetailsTime := time.Now()
	detailsCount := 0
	if details.GetFields() != nil {
		detailsCount = len(details.GetFields())
	}

	if !hasError {
		saveIndexedHash()
	}

	metrics.SharedClient.RecordEvent(metrics.IndexEvent{
		ObjectId:                info.Id,
		IndexLinksTimeMs:        indexLinksTime.Sub(indexSetTime).Milliseconds(),
		IndexDetailsTimeMs:      indexDetailsTime.Sub(indexLinksTime).Milliseconds(),
		IndexSetRelationsTimeMs: indexSetTime.Sub(startTime).Milliseconds(),
		DetailsCount:            detailsCount,
	})

	return nil
}

func (i *indexer) indexLinkedFiles(ctx context.Context, space smartblock.Space, fileHashes []string) {
	if len(fileHashes) == 0 {
		return
	}
	existingIDs, err := i.store.HasIDs(fileHashes...)
	if err != nil {
		log.Errorf("failed to get existing file ids : %s", err)
	}
	newIDs := slice.Difference(fileHashes, existingIDs)
	for _, id := range newIDs {
		go func(id string) {
			// Deduplicate
			_, ok := i.indexedFiles.LoadOrStore(id, struct{}{})
			if ok {
				return
			}
			bindErr := i.storageService.BindSpaceID(space.Id(), id)
			if bindErr != nil {
				log.Error("failed to bind space id", zap.Error(bindErr), zap.String("id", id))
				return
			}
			// file's hash is id
			idxErr := i.reindexDoc(ctx, space, id)
			if idxErr != nil && !errors.Is(idxErr, domain.ErrFileNotFound) {
				log.With("id", id).Errorf("failed to reindex file: %s", idxErr)
			}
			idxErr = i.store.AddToIndexQueue(id)
			if idxErr != nil {
				log.With("id", id).Error(idxErr)
			}
		}(id)
	}
}

func headsHash(heads []string) string {
	if len(heads) == 0 {
		return ""
	}
	slices.Sort(heads)

	sum := sha256.Sum256([]byte(strings.Join(heads, ",")))
	return fmt.Sprintf("%x", sum)
}
