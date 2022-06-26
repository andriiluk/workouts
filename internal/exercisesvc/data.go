package exercisesvc

import "github.com/andriiluk/workouts/internal"

type PostExerciseRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Muscles     []string `json:"muscles,omitempty"`
}

type PostExerciseResponse struct {
	ID  int    `json:"id,omitempty"`
	Err string `json:"err,omitempty"`
}

type GetExerciseRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetExerciseResponse struct {
	Err      error              `json:"err,omitempty"`
	Exercise *internal.Exercise `json:"exercise,omitempty"`
}

type PutExerciseRequest struct {
	ID       int                `json:"id,omitempty"`
	Exercise *internal.Exercise `json:"exercise,omitempty"`
}

type DeleteExerciseRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetExercisesByMusclesRequest struct {
	Muscles []string `json:"muscles,omitempty"`
}

type GetExercisesByMusclesResponse struct {
	Err       string               `json:"err,omitempty"`
	Exercises []*internal.Exercise `json:"exercises,omitempty"`
}

type SearchExercisesByTagsRequest struct {
	Tags []string `json:"tags,omitempty"`
}

type SearchExercisesByTagsResponse struct {
	Err       error                `json:"err,omitempty"`
	Exercises []*internal.Exercise `json:"exercises,omitempty"`
}

type DefaultResponse struct {
	Err error `json:"err,omitempty"`
}
