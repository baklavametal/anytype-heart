package space

import (
	"context"
	"fmt"

	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/space/internal/spacecontroller"
	"github.com/anyproto/anytype-heart/space/internal/spaceprocess/loader"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

type controllerWaiter struct {
	wait chan struct{}
	err  error
}

func (s *service) startStatus(ctx context.Context, spaceId string, status spaceinfo.AccountStatus) (ctrl spacecontroller.SpaceController, err error) {
	s.mu.Lock()
	if ctrl, ok := s.spaceControllers[spaceId]; ok {
		s.mu.Unlock()
		return ctrl, nil
	}
	if w, ok := s.waiting[spaceId]; ok {
		s.mu.Unlock()
		select {
		case <-w.wait:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
		s.mu.Lock()
		err := s.waiting[spaceId].err
		if err != nil {
			s.mu.Unlock()
			return nil, err
		}
		ctrl := s.spaceControllers[spaceId]
		s.mu.Unlock()
		return ctrl, nil
	}
	wait := make(chan struct{})
	s.waiting[spaceId] = controllerWaiter{
		wait: wait,
	}
	s.mu.Unlock()
	ctrl, err = s.factory.NewShareableSpace(ctx, spaceId, status)
	s.mu.Lock()
	close(wait)
	if err != nil {
		s.waiting[spaceId] = controllerWaiter{
			wait: wait,
			err:  err,
		}
		s.mu.Unlock()
		return nil, err
	}
	s.spaceControllers[spaceId] = ctrl
	s.mu.Unlock()
	return ctrl, nil
}

func (s *service) waitLoad(ctx context.Context, ctrl spacecontroller.SpaceController) (sp clientspace.Space, err error) {
	if ld, ok := ctrl.Current().(loader.LoadWaiter); ok {
		return ld.WaitLoad(ctx)
	}
	return nil, fmt.Errorf("failed to load space, mode is %d", ctrl.Mode())
}

func (s *service) loadPersonalSpace(ctx context.Context) (err error) {
	s.mu.Lock()
	wait := make(chan struct{})
	s.waiting[s.personalSpaceId] = controllerWaiter{
		wait: wait,
	}
	s.mu.Unlock()
	ctrl, err := s.factory.NewPersonalSpace(ctx)
	if err != nil {
		return
	}
	_, err = ctrl.Current().(loader.LoadWaiter).WaitLoad(ctx)
	s.mu.Lock()
	defer s.mu.Unlock()
	if err != nil {
		return err
	}
	close(wait)
	s.spaceControllers[s.personalSpaceId] = ctrl
	return
}
