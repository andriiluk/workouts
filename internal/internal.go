package internal

import (
	"context"
	"errors"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrBadResponse           = errors.New("bad response")
	ErrInternalService       = errors.New("internal service error")
	ErrDecodeRequest         = errors.New("error decoding request")
	ErrMissingRequiredParams = errors.New("missing required parameters")
)

const (
	ExerciseSvcTagName = "exercise_svc"
	MuscleSvcTagName   = "muscle_svc"
	WorkoutSvcTagName  = "workout_svc"
	RoutineSvcTagName  = "routine_svc"
)

type StorageConstraints interface {
	Exercise | Muscle | Workout
}

type Storage[T StorageConstraints] interface {
	InsertOrUpdate(ctx context.Context, t *T) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*T, error)
	Search(ctx context.Context, p *Params) ([]*T, error)
}

type Tag struct {
	Name        string `json:"name,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
}

type Params struct {
	ID      int
	Name    string
	Tags    []string
	Muscles []string
}

type Muscle struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Exercise struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Muscles     []string `json:"muscles,omitempty"`
}

type Workout struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Tags        []*Tag      `json:"tags,omitempty"`
	Exercise    []*Exercise `json:"exercise,omitempty"`
}

type Routine struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Tags        []*Tag     `json:"tags,omitempty"`
	Workouts    []*Workout `json:"workouts,omitempty"`
}
