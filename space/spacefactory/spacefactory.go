package spacefactory

import (
	"context"
	"errors"
	"fmt"

	"github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"

	"github.com/anyproto/anytype-heart/core/block/object/objectcache"
	"github.com/anyproto/anytype-heart/space/clientspace"
	dependencies "github.com/anyproto/anytype-heart/space/internal/components/dependencies"
	"github.com/anyproto/anytype-heart/space/internal/marketplacespace"
	"github.com/anyproto/anytype-heart/space/internal/personalspace"
	"github.com/anyproto/anytype-heart/space/internal/shareablespace"
	"github.com/anyproto/anytype-heart/space/internal/spacecontroller"
	"github.com/anyproto/anytype-heart/space/internal/techspace"
	"github.com/anyproto/anytype-heart/space/spacecore"
	"github.com/anyproto/anytype-heart/space/spacecore/storage"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

type SpaceFactory interface {
	app.Component
	CreatePersonalSpace(ctx context.Context) (sp spacecontroller.SpaceController, err error)
	NewPersonalSpace(ctx context.Context) (spacecontroller.SpaceController, error)
	CreateShareableSpace(ctx context.Context, id string) (sp spacecontroller.SpaceController, err error)
	NewShareableSpace(ctx context.Context, id string, status spaceinfo.AccountStatus) (spacecontroller.SpaceController, error)
	CreateMarketplaceSpace(ctx context.Context) (sp spacecontroller.SpaceController, err error)
	CreateAndSetTechSpace(ctx context.Context) (*clientspace.TechSpace, error)
}

const CName = "client.space.spacefactory"

type spaceFactory struct {
	app             *app.App
	spaceCore       spacecore.SpaceCoreService
	techSpace       techspace.TechSpace
	accountService  accountservice.Service
	objectFactory   objectcache.ObjectFactory
	indexer         dependencies.SpaceIndexer
	installer       dependencies.BundledObjectsInstaller
	storageService  storage.ClientStorage
	personalSpaceId string
}

func New() SpaceFactory {
	return &spaceFactory{}
}

func (s *spaceFactory) Init(a *app.App) (err error) {
	s.app = a
	s.spaceCore = app.MustComponent[spacecore.SpaceCoreService](a)
	s.accountService = app.MustComponent[accountservice.Service](a)
	s.objectFactory = app.MustComponent[objectcache.ObjectFactory](a)
	s.indexer = app.MustComponent[dependencies.SpaceIndexer](a)
	s.installer = app.MustComponent[dependencies.BundledObjectsInstaller](a)
	s.storageService = app.MustComponent[storage.ClientStorage](a)
	s.personalSpaceId, err = s.spaceCore.DeriveID(context.Background(), spacecore.SpaceType)
	if err != nil {
		return
	}
	return
}

func (s *spaceFactory) CreatePersonalSpace(ctx context.Context) (sp spacecontroller.SpaceController, err error) {
	coreSpace, err := s.spaceCore.Derive(ctx, spacecore.SpaceType)
	if err != nil {
		return
	}
	err = s.storageService.MarkSpaceCreated(coreSpace.Id())
	if err != nil {
		return
	}
	if err := s.techSpace.SpaceViewCreate(ctx, coreSpace.Id(), true); err != nil {
		if errors.Is(err, techspace.ErrSpaceViewExists) {
			return s.NewPersonalSpace(ctx)
		}
		return nil, err
	}
	ctrl := personalspace.NewSpaceController(coreSpace.Id(), s.app)
	err = ctrl.Start(ctx)
	return ctrl, err
}

func (s *spaceFactory) NewPersonalSpace(ctx context.Context) (ctrl spacecontroller.SpaceController, err error) {
	coreSpace, err := s.spaceCore.Derive(ctx, spacecore.SpaceType)
	if err != nil {
		return nil, err
	}
	ctrl = personalspace.NewSpaceController(coreSpace.Id(), s.app)
	err = ctrl.Start(ctx)
	return ctrl, err
}

func (s *spaceFactory) CreateAndSetTechSpace(ctx context.Context) (*clientspace.TechSpace, error) {
	techSpace := techspace.New()
	techCoreSpace, err := s.spaceCore.Derive(ctx, spacecore.TechSpaceType)
	if err != nil {
		return nil, fmt.Errorf("derive tech space: %w", err)
	}
	deps := clientspace.TechSpaceDeps{
		CommonSpace:     techCoreSpace,
		ObjectFactory:   s.objectFactory,
		AccountService:  s.accountService,
		PersonalSpaceId: s.personalSpaceId,
		Indexer:         s.indexer,
		Installer:       s.installer,
		TechSpace:       techSpace,
	}
	ts := clientspace.NewTechSpace(deps)
	err = ts.Run(techCoreSpace, ts.Cache)
	if err != nil {
		return nil, fmt.Errorf("run tech space: %w", err)
	}

	s.techSpace = ts
	s.app = s.app.ChildApp()
	s.app.Register(s.techSpace)
	return ts, nil
}

func (s *spaceFactory) NewShareableSpace(ctx context.Context, id string, status spaceinfo.AccountStatus) (spacecontroller.SpaceController, error) {
	ctrl, err := shareablespace.NewSpaceController(id, status, s.app)
	if err != nil {
		return nil, err
	}
	err = ctrl.Start(ctx)
	return ctrl, err
}

func (s *spaceFactory) CreateShareableSpace(ctx context.Context, id string) (sp spacecontroller.SpaceController, err error) {
	err = s.storageService.MarkSpaceCreated(id)
	if err != nil {
		return
	}
	if err := s.techSpace.SpaceViewCreate(ctx, id, true); err != nil {
		return nil, err
	}
	ctrl, err := shareablespace.NewSpaceController(id, spaceinfo.AccountStatusUnknown, s.app)
	if err != nil {
		return nil, err
	}
	err = ctrl.Start(ctx)
	return ctrl, err
}

func (s *spaceFactory) CreateMarketplaceSpace(ctx context.Context) (sp spacecontroller.SpaceController, err error) {
	ctrl := marketplacespace.NewSpaceController(s.app, s.personalSpaceId)
	err = ctrl.Start(ctx)
	return ctrl, err
}

func (s *spaceFactory) Name() (name string) {
	return CName
}
