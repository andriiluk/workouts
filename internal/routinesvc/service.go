package routinesvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

type Service interface {
	AddRoutine(ctx context.Context, m *internal.Routine) (int, error)
	PutRoutine(ctx context.Context, m *internal.Routine) error
	DeleteRoutineByName(ctx context.Context, name string) error
	DeleteRoutineByID(ctx context.Context, id int) error
	GetRoutineByID(ctx context.Context, id int) (*internal.Routine, error)
	GetRoutineByName(ctx context.Context, name string) (*internal.Routine, error)
	GetRoutinesByTags(ctx context.Context, tags ...*internal.Tag) ([]*internal.Routine, error)
}

type service struct {
	store internal.Storage[internal.Muscle]
}

func NewService(store internal.Storage[internal.Muscle]) *service {
	return &service{
		store: store,
	}
}

func (s *service) AddRoutine(ctx context.Context, m *internal.Routine) (int, error) {
	panic("not implemented yet")
}

func (s *service) PutRoutine(ctx context.Context, m *internal.Routine) error {
	panic("not implemented yet")
}

func (s *service) DeleteRoutineByName(ctx context.Context, name string) error {
	// todo: delete Routine
	panic("not implemented yet")
}

func (s *service) DeleteRoutineByID(ctx context.Context, id int) error {
	// todo: delete Routine
	panic("not implemented yet")
}

func (s *service) GetRoutineByID(ctx context.Context, id int) ([]*internal.Routine, error) {
	// todo: get Routine
	panic("not implemented yet")
}

func (s *service) GetRoutineByName(ctx context.Context, name string) ([]*internal.Routine, error) {
	// todo: get Routine
	panic("not implemented yet")
}

func (s *service) GetRoutinesByTags(ctx context.Context, tags ...internal.Tag) ([]*internal.Routine, error) {
	// todo: get Routine
	panic("not implemented yet")
}
