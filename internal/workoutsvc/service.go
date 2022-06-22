package workoutsvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

type Service interface {
	AddWorkout(ctx context.Context, m *internal.Workout) (int, error)
	PutWorkout(ctx context.Context, m *internal.Workout) error
	DeleteWorkoutByName(ctx context.Context, name string) error
	DeleteWorkoutByID(ctx context.Context, id int) error
	GetWorkoutByID(ctx context.Context, id int) (*internal.Workout, error)
	GetWorkoutByName(ctx context.Context, name string) (*internal.Workout, error)
	GetWorkoutsByTags(ctx context.Context, tags ...*internal.Tag) ([]*internal.Workout, error)
}

type service struct {
	store internal.Storage[internal.Muscle]
}

func NewService(store internal.Storage[internal.Muscle]) *service {
	return &service{
		store: store,
	}
}

func (s *service) AddWorkout(ctx context.Context, m *internal.Workout) (int, error) {
	panic("not implemented yet")
}

func (s *service) PutWorkout(ctx context.Context, m *internal.Workout) error {
	panic("not implemented yet")
}

func (s *service) DeleteWorkoutByName(ctx context.Context, name string) error {
	// todo: delete Workout
	panic("not implemented yet")
}

func (s *service) DeleteWorkoutByID(ctx context.Context, id int) error {
	// todo: delete Workout
	panic("not implemented yet")
}

func (s *service) GetWorkoutByID(ctx context.Context, id int) ([]*internal.Workout, error) {
	// todo: get Workout
	panic("not implemented yet")
}

func (s *service) GetWorkoutByName(ctx context.Context, name string) ([]*internal.Workout, error) {
	// todo: get Workout
	panic("not implemented yet")
}

func (s *service) GetWorkoutsByTags(ctx context.Context, tags ...internal.Tag) ([]*internal.Workout, error) {
	// todo: get Workout
	panic("not implemented yet")
}
