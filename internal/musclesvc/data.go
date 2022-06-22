package musclesvc

import "github.com/andriiluk/workouts/internal"

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
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetMuscleResponse struct {
	Err    error            `json:"err,omitempty"`
	Muscle *internal.Muscle `json:"muscle,omitempty"`
}

type PutMuscleRequest struct {
	ID     int              `json:"id,omitempty"`
	Muscle *internal.Muscle `json:"muscle,omitempty"`
}

type DeleteMuscleRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
