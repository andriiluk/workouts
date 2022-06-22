package exercisesvc

import (
	"context"
	"strings"

	"github.com/andriiluk/workouts/internal"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddExercise     endpoint.Endpoint
	PutExercise     endpoint.Endpoint
	DeleteExercise  endpoint.Endpoint
	GetExercise     endpoint.Endpoint
	SearchExercises endpoint.Endpoint
}

func MakeEndpoints(svc Service) *Endpoints {
	return &Endpoints{
		AddExercise:     makeAddExerciseEndpoint(svc),
		PutExercise:     makePutExerciseEndpoint(svc),
		DeleteExercise:  makeDeleteExerciseEndpoint(svc),
		GetExercise:     makeGetExerciseEndpoint(svc),
		SearchExercises: makeSearchExerciseEndpoint(svc),
	}
}

func makeAddExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addExerciseRequest)

		id, err := svc.AddExercise(ctx, &internal.Exercise{
			Name:        req.Name,
			Description: req.Description,
			Tags:        req.Tags,
			Muscles:     req.Muscles,
		})

		return addExerciseResponse{
			Err: err,
			ID:  id,
		}, nil
	}
}

func makePutExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(putExerciseRequest)

		err = svc.PutExercise(ctx, &internal.Exercise{
			ID:          req.ID,
			Name:        req.Exercise.Name,
			Description: req.Exercise.Description,
			Tags:        req.Exercise.Tags,
			Muscles:     req.Exercise.Muscles,
		})

		return defaultResponse{
			Err: err,
		}, nil
	}
}

func makeDeleteExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteExerciseRequest)
		resp := defaultResponse{}

		switch {
		case req.ID != 0:
			resp.Err = svc.DeleteExerciseByID(ctx, req.ID)
		case strings.TrimSpace(req.Name) != "":
			resp.Err = svc.DeleteExerciseByName(ctx, req.Name)
		}

		return resp, nil
	}
}

func makeGetExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getExerciseRequest)
		resp := getExerciseResponse{}

		switch {
		case req.ID != 0:
			resp.Exercise, resp.Err = svc.GetExerciseByID(ctx, req.ID)
		case strings.TrimSpace(req.Name) != "":
			resp.Exercise, resp.Err = svc.GetExerciseByName(ctx, req.Name)
		}

		return resp, nil
	}
}

func makeSearchExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(searchExercisesByTagsRequest)
		resp := searchExercisesByTagsResponse{}

		resp.Exercises, resp.Err = svc.GetExercisesByTags(ctx, req.Tags...)

		return resp, nil
	}
}

type addExerciseRequest struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Tags        []*internal.Tag    `json:"tags,omitempty"`
	Muscles     []*internal.Muscle `json:"muscles,omitempty"`
}

type addExerciseResponse struct {
	Err error
	ID  int
}

type putExerciseRequest struct {
	ID       int
	Exercise internal.Exercise
}

type deleteExerciseRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type getExerciseRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type getExerciseResponse struct {
	Err      error              `json:"err,omitempty"`
	Exercise *internal.Exercise `json:"exercises,omitempty"`
}

type searchExercisesByTagsRequest struct {
	Tags []*internal.Tag `json:"tags,omitempty"`
}

type searchExercisesByTagsResponse struct {
	Err       error                `json:"err,omitempty"`
	Exercises []*internal.Exercise `json:"exercises,omitempty"`
}

type defaultResponse struct {
	Err error `json:"err,omitempty"`
}
