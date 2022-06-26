package exercisesvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

const SvcName = "ExerciseSvc"

type Service interface {
	AddExercise(ctx context.Context, ex *internal.Exercise) (int, error)
	PutExercise(ctx context.Context, ex *internal.Exercise) error
	DeleteExerciseByName(ctx context.Context, name string) error
	DeleteExerciseByID(ctx context.Context, id int) error
	GetExerciseByID(ctx context.Context, id int) (*internal.Exercise, error)
	GetExerciseByName(ctx context.Context, name string) (*internal.Exercise, error)
	GetExercisesByTags(ctx context.Context, tags ...string) ([]*internal.Exercise, error)
	GetExercisesByMuscles(ctx context.Context, muscleNames ...string) ([]*internal.Exercise, error)
}

type service struct {
	store internal.Storage[internal.Exercise]
}

func NewService(store internal.Storage[internal.Exercise]) Service {
	return &service{
		store: store,
	}
}

func (s *service) AddExercise(ctx context.Context, m *internal.Exercise) (int, error) {
	if err := s.store.InsertOrUpdate(ctx, m); err != nil {
		return 0, err
	}

	return m.ID, nil
}

func (s *service) GetExerciseByID(ctx context.Context, id int) (*internal.Exercise, error) {
	return s.store.Get(ctx, id)
}

func (s *service) PutExercise(ctx context.Context, m *internal.Exercise) error {
	return s.store.InsertOrUpdate(ctx, m)
}

func (s *service) DeleteExerciseByID(ctx context.Context, id int) error {
	return s.store.Delete(ctx, id)
}

func (s *service) DeleteExerciseByName(ctx context.Context, name string) error {
	// todo: delete Exercise
	panic("not implemented yet")
}

func (s *service) GetExerciseByName(ctx context.Context, name string) (*internal.Exercise, error) {
	// todo: get Exercise
	panic("not implemented yet")
}

func (s *service) GetExercisesByTags(ctx context.Context, tags ...string) ([]*internal.Exercise, error) {
	return s.store.Search(ctx, &internal.Params{
		Tags: tags,
	})
}

func (s *service) GetExercisesByMuscles(ctx context.Context, muscleNames ...string) ([]*internal.Exercise, error) {
	// todo: get exercise by muscle
	panic("not implemented yet")
}
