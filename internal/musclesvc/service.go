package musclesvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

const SvcName = "MuscleSvc"

type Service interface {
	AddMuscle(ctx context.Context, m *internal.Muscle) (int, error)
	PutMuscle(ctx context.Context, m *internal.Muscle) error
	DeleteMuscleByName(ctx context.Context, name string) error
	DeleteMuscleByID(ctx context.Context, id int) error
	GetMuscleByID(ctx context.Context, id int) (*internal.Muscle, error)
	GetMuscleByName(ctx context.Context, name string) (*internal.Muscle, error)
	GetMusclesByTags(ctx context.Context, tags ...string) ([]*internal.Muscle, error)
}

type service struct {
	store internal.Storage[internal.Muscle]
}

func NewService(store internal.Storage[internal.Muscle]) Service {
	return &service{
		store: store,
	}
}

func (s *service) AddMuscle(ctx context.Context, m *internal.Muscle) (int, error) {
	if err := s.store.InsertOrUpdate(ctx, m); err != nil {
		return 0, err
	}

	return m.ID, nil
}

func (s *service) GetMuscleByID(ctx context.Context, id int) (*internal.Muscle, error) {
	return s.store.Get(ctx, id)
}

func (s *service) PutMuscle(ctx context.Context, m *internal.Muscle) error {
	return s.store.InsertOrUpdate(ctx, m)
}

func (s *service) DeleteMuscleByID(ctx context.Context, id int) error {
	return s.store.Delete(ctx, id)
}

func (s *service) DeleteMuscleByName(ctx context.Context, name string) error {
	// todo: delete Muscle
	panic("not implemented yet")
}

func (s *service) GetMuscleByName(ctx context.Context, name string) (*internal.Muscle, error) {
	// todo: get Muscle
	panic("not implemented yet")
}

func (s *service) GetMusclesByTags(ctx context.Context, tags ...string) ([]*internal.Muscle, error) {
	return s.store.Search(ctx, &internal.Params{
		Tags: tags,
	})
}
