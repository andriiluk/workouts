package exercisesvc

import (
	"context"

	"fmt"

	"strings"

	"github.com/go-kit/kit/endpoint"

	"github.com/andriiluk/workouts/internal"
)

type Endpoints struct {
	PostExerciseEndpoint          endpoint.Endpoint
	PutExerciseEndpoint           endpoint.Endpoint
	DeleteExerciseEndpoint        endpoint.Endpoint
	GetExerciseEndpoint           endpoint.Endpoint
	SearchExercisesEndpoint       endpoint.Endpoint
	GetExercisesByMusclesEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service) *Endpoints {
	return &Endpoints{
		PostExerciseEndpoint:          makePostExerciseEndpoint(svc),
		PutExerciseEndpoint:           makePutExerciseEndpoint(svc),
		DeleteExerciseEndpoint:        makeDeleteExerciseEndpoint(svc),
		GetExerciseEndpoint:           makeGetExerciseEndpoint(svc),
		SearchExercisesEndpoint:       makeSearchExerciseEndpoint(svc),
		GetExercisesByMusclesEndpoint: makeGetExercisesByMusclesEndpoint(svc),
	}
}

func makePostExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var resp PostExerciseResponse

		req, ok := request.(PostExerciseRequest)
		if !ok {
			return nil, fmt.Errorf("[%w] request exptected to be PostExerciseRequest", internal.ErrBadRequest)
		}

		if strings.TrimSpace(req.Name) == "" {
			return PostExerciseResponse{
				Err: fmt.Sprintf("[%s]: 'Name'", internal.ErrMissingRequiredParams),
			}, nil
		}

		id, err := svc.AddExercise(ctx, &internal.Exercise{
			Name:        req.Name,
			Description: req.Description,
			Tags:        req.Tags,
			Muscles:     req.Muscles,
		})

		if err != nil {
			resp.Err = err.Error()
		}

		resp.ID = id

		return resp, nil
	}
}

func makePutExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(PutExerciseRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		err = svc.PutExercise(ctx, &internal.Exercise{
			ID:          req.ID,
			Name:        req.Exercise.Name,
			Description: req.Exercise.Description,
			Tags:        req.Exercise.Tags,
			Muscles:     req.Exercise.Muscles,
		})

		return DefaultResponse{
			Err: err,
		}, nil
	}
}

func makeDeleteExerciseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteExerciseRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		resp := DefaultResponse{}

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
		req, ok := request.(GetExerciseRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		resp := GetExerciseResponse{}

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
		req, ok := request.(SearchExercisesByTagsRequest)
		if !ok {
			return nil, fmt.Errorf("[%w]: request expected to be SearchExercisesByTagsRequest", internal.ErrBadRequest)
		}

		resp := SearchExercisesByTagsResponse{}

		resp.Exercises, resp.Err = svc.GetExercisesByTags(ctx, req.Tags...)

		return resp, nil
	}
}

func makeGetExercisesByMusclesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetExercisesByMusclesRequest)
		if !ok {
			return nil, fmt.Errorf("[%w]: request expected to be SearchExercisesByTagsRequest", internal.ErrBadRequest)
		}

		exercises, err := svc.GetExercisesByMuscles(ctx, req.Muscles...)
		if err != nil {
			return GetExercisesByMusclesResponse{
				Err: err.Error(),
			}, nil
		}

		return GetExercisesByMusclesResponse{
			Exercises: exercises,
		}, nil
	}
}
