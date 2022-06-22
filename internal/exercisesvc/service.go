package exercisesvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

type Service interface {
	AddExercise(ctx context.Context, m *internal.Exercise) (int, error)
	PutExercise(ctx context.Context, m *internal.Exercise) error
	DeleteExerciseByName(ctx context.Context, name string) error
	DeleteExerciseByID(ctx context.Context, id int) error
	GetExerciseByID(ctx context.Context, id int) (*internal.Exercise, error)
	GetExerciseByName(ctx context.Context, name string) (*internal.Exercise, error)
	GetExercisesByTags(ctx context.Context, tags ...*internal.Tag) ([]*internal.Exercise, error)
	AddMusclesToExercise(ctx context.Context, id int, muscles ...*internal.Muscle)
}

type service struct {
	store internal.Storage[internal.Muscle]
}

func NewService(store internal.Storage[internal.Muscle]) *service {
	return &service{
		store: store,
	}
}

func (s *service) AddExercise(ctx context.Context, m *internal.Exercise) (int, error) {
	panic("not implemented yet")
}

func (s *service) PutExercise(ctx context.Context, m *internal.Exercise) error {
	panic("not implemented yet")
}

func (s *service) DeleteExerciseByName(ctx context.Context, name string) error {
	// todo: delete exercise
	panic("not implemented yet")
}

func (s *service) DeleteExerciseByID(ctx context.Context, id int) error {
	// todo: delete exercise
	panic("not implemented yet")
}

func (s *service) GetExerciseByID(ctx context.Context, id int) ([]*internal.Exercise, error) {
	// todo: get exercise
	panic("not implemented yet")
}

func (s *service) GetExerciseByName(ctx context.Context, name string) ([]*internal.Exercise, error) {
	// todo: get exercise
	panic("not implemented yet")
}

func (s *service) GetExercisesByTags(ctx context.Context, tags ...internal.Tag) ([]*internal.Exercise, error) {
	// todo: get exercise
	panic("not implemented yet")
}
