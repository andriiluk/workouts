package musclesvc

import (
	"context"

	"fmt"

	"strings"

	"github.com/go-kit/kit/endpoint"

	"github.com/andriiluk/workouts/internal"
)

type Endpoints struct {
	PostMuscleEndpoint    endpoint.Endpoint
	PutMuscleEndpoint     endpoint.Endpoint
	DeleteMuscleEndpoint  endpoint.Endpoint
	GetMuscleEndpoint     endpoint.Endpoint
	SearchMusclesEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service) *Endpoints {
	return &Endpoints{
		PostMuscleEndpoint:    makePostMuscleEndpoint(svc),
		PutMuscleEndpoint:     makePutMuscleEndpoint(svc),
		DeleteMuscleEndpoint:  makeDeleteMuscleEndpoint(svc),
		GetMuscleEndpoint:     makeGetMuscleEndpoint(svc),
		SearchMusclesEndpoint: makeSearchMuscleEndpoint(svc),
	}
}

func makePostMuscleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var resp PostMuscleResponse

		req, ok := request.(PostMuscleRequest)
		if !ok {
			return nil, fmt.Errorf("[%w] request exptected to be PostMuscleRequest", internal.ErrBadRequest)
		}

		if strings.TrimSpace(req.Name) == "" {
			return PostMuscleResponse{
				Err: fmt.Sprintf("[%s]: 'Name'", internal.ErrMissingRequiredParams),
			}, nil
		}

		id, err := svc.AddMuscle(ctx, &internal.Muscle{
			Name:        req.Name,
			Description: req.Description,
			Tags:        req.Tags,
		})

		if err != nil {
			resp.Err = err.Error()
		}

		resp.ID = id

		return resp, nil
	}
}

func makePutMuscleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(PutMuscleRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		err = svc.PutMuscle(ctx, &internal.Muscle{
			ID:          req.ID,
			Name:        req.Muscle.Name,
			Description: req.Muscle.Description,
			Tags:        req.Muscle.Tags,
		})

		return DefaultResponse{
			Err: err,
		}, nil
	}
}

func makeDeleteMuscleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteMuscleRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		resp := DefaultResponse{}

		switch {
		case req.ID != 0:
			resp.Err = svc.DeleteMuscleByID(ctx, req.ID)
		case strings.TrimSpace(req.Name) != "":
			resp.Err = svc.DeleteMuscleByName(ctx, req.Name)
		}

		return resp, nil
	}
}

func makeGetMuscleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetMuscleRequest)
		if !ok {
			return nil, internal.ErrBadRequest
		}

		resp := GetMuscleResponse{}

		var err error

		switch {
		case req.ID != 0:
			resp.Muscle, err = svc.GetMuscleByID(ctx, req.ID)
		case strings.TrimSpace(req.Name) != "":
			resp.Muscle, err = svc.GetMuscleByName(ctx, req.Name)
		}

		if err != nil {
			resp.Err = err.Error()
		}

		return resp, nil
	}
}

func makeSearchMuscleEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(SearchMusclesByTagsRequest)
		if !ok {
			return nil, fmt.Errorf("[%w]: request expected to be SearchMusclesByTagsRequest", internal.ErrBadRequest)
		}

		resp := SearchMusclesByTagsResponse{}

		resp.Muscles, resp.Err = svc.GetMusclesByTags(ctx, req.Tags...)

		return resp, nil
	}
}
