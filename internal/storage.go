package internal

import "context"

type StorageConstraints interface {
	Exercise | Muscle | Workout
}

type Storage[T StorageConstraints] interface {
	InsertOrUpdate(ctx context.Context, t *T) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*T, error)
	Search(ctx context.Context, p *Params) ([]*T, error)
}
