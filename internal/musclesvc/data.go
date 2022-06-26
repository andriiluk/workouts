package musclesvc

import (
	"github.com/andriiluk/workouts/internal"
)

type PostMuscleRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type PostMuscleResponse struct {
	Err string `json:"err,omitempty"`
	ID  int    `json:"id,omitempty"`
}

type GetMuscleRequest struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type GetMuscleResponse struct {
	Err    string           `json:"err,omitempty"`
	Muscle *internal.Muscle `json:"muscle,omitempty"`
}

type PutMuscleRequest struct {
	Muscle *internal.Muscle `json:"muscle,omitempty"`
	ID     int              `json:"id,omitempty"`
}

type DeleteMuscleRequest struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type SearchMusclesByTagsRequest struct {
	Tags []string `json:"tags,omitempty"`
}

type SearchMusclesByTagsResponse struct {
	Err     error              `json:"err,omitempty"`
	Muscles []*internal.Muscle `json:"muscles,omitempty"`
}

type DefaultResponse struct {
	Err error `json:"err,omitempty"`
}
