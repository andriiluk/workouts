package service

import (
	"context"

	"github.com/andriiluk/workouts/internal"
)

type Service interface {
	AddMuscle(ctx context.Context, m internal.Muscle) error
	DeleteMuscle(ctx context.Context, p internal.Params) error
	GetMuscles(ctx context.Context, p internal.Params) []internal.Muscle

	AddExercise(ctx context.Context, m internal.Exercise) error
	DeleteExercise(ctx context.Context, p internal.Params) error
	GetExercises(ctx context.Context, p internal.Params) []internal.Exercise

	AddWorkout(ctx context.Context, m internal.Workout) error
	DeleteWorkout(ctx context.Context, p internal.Params) error
	GetWorkouts(ctx context.Context, p internal.Params) []internal.Workout
}

type StorageConstraints interface {
	internal.Exercise | internal.Muscle | internal.Workout
}

type Storage[T StorageConstraints] interface {
	InsertOrUpdate(ctx context.Context, t T) error
	Delete(ctx context.Context, p internal.Params) error
	Get(ctx context.Context, p internal.Params) []T
}

type service struct {
	muscleStore   Storage[internal.Muscle]
	exerciseStore Storage[internal.Exercise]
	workoutStore  Storage[internal.Workout]
}

func NewService(muscleStore Storage[internal.Muscle], exerciseStore Storage[internal.Exercise], workoutStore Storage[internal.Workout]) *service {
	return &service{
		muscleStore:   muscleStore,
		exerciseStore: exerciseStore,
		workoutStore:  workoutStore,
	}
}

func (s *service) AddMuscle(ctx context.Context, m internal.Muscle) error {
	panic("not implemented yet")
}

func (s *service) DeleteMuscle(ctx context.Context, p internal.Params) error {
	panic("not implemented yet")
}

func (s *service) GetMuscles(ctx context.Context, p internal.Params) []internal.Muscle {
	panic("not implemented yet")
}

func (s *service) AddExercise(ctx context.Context, m internal.Exercise) error {
	panic("not implemented yet")
}

func (s *service) DeleteExercise(ctx context.Context, p internal.Params) error {
	panic("not implemented yet")
}

func (s *service) GetExercises(ctx context.Context, p internal.Params) []internal.Exercise {
	panic("not implemented yet")
}

func (s *service) AddWorkout(ctx context.Context, m internal.Workout) error {
	panic("not implemented yet")
}

func (s *service) DeleteWorkout(ctx context.Context, p internal.Params) error {
	panic("not implemented yet")
}

func (s *service) GetWorkouts(ctx context.Context, p internal.Params) []internal.Workout {
	panic("not implemented yet")
}
