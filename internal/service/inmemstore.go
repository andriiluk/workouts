package service

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

type inMemStore struct {
}

func NewInMemstore() *inMemStore {
	return &inMemStore{}
}

func (s inMemStore) InsertOrUpdateMuscle(ctx context.Context, m internal.Muscle) error {
	// TODO: muscle in-memory store
	panic("not implemented yet")
}

func (s inMemStore) DeleteMuscle(ctx context.Context, p internal.Params) error {
	// TODO: muscle in-memory store
	panic("not implemented yet")
}

func (s inMemStore) GetMuscles(ctx context.Context, p internal.Params) []internal.Muscle {
	// TODO: muscle in-memory store
	panic("not implemented yet")
}
