package routinesvc

// import (
// 	"context"
// 	"strings"

// 	"github.com/andriiluk/workouts/internal"
// 	"github.com/go-kit/kit/endpoint"
// )

// type Endpoints struct {
// 	AddRoutine     endpoint.Endpoint
// 	PutRoutine     endpoint.Endpoint
// 	DeleteRoutine  endpoint.Endpoint
// 	GetRoutine     endpoint.Endpoint
// 	SearchRoutines endpoint.Endpoint
// }

// func MakeEndpoints(svc Service) *Endpoints {
// 	return &Endpoints{
// 		AddRoutine:     makeAddRoutineEndpoint(svc),
// 		PutRoutine:     makePutRoutineEndpoint(svc),
// 		DeleteRoutine:  makeDeleteRoutineEndpoint(svc),
// 		GetRoutine:     makeGetRoutineEndpoint(svc),
// 		SearchRoutines: makeSearchRoutineEndpoint(svc),
// 	}
// }

// func makeAddRoutineEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(addRoutineRequest)

// 		id, err := svc.AddRoutine(ctx, &internal.Routine{
// 			Name:        req.Name,
// 			Description: req.Description,
// 			Tags:        req.Tags,
// 		})

// 		return addRoutineResponse{
// 			Err: err,
// 			ID:  id,
// 		}, nil
// 	}
// }

// func makePutRoutineEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(putRoutineRequest)

// 		err = svc.PutRoutine(ctx, &internal.Routine{
// 			ID:          req.ID,
// 			Name:        req.Routine.Name,
// 			Description: req.Routine.Description,
// 			Tags:        req.Routine.Tags,
// 		})

// 		return defaultResponse{
// 			Err: err,
// 		}, nil
// 	}
// }

// func makeDeleteRoutineEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(deleteRoutineRequest)
// 		resp := defaultResponse{}

// 		switch {
// 		case req.ID != 0:
// 			resp.Err = svc.DeleteRoutineByID(ctx, req.ID)
// 		case strings.TrimSpace(req.Name) != "":
// 			resp.Err = svc.DeleteRoutineByName(ctx, req.Name)
// 		}

// 		return resp, nil
// 	}
// }

// func makeGetRoutineEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(getRoutineRequest)
// 		resp := getRoutineResponse{}

// 		switch {
// 		case req.ID != 0:
// 			resp.Routine, resp.Err = svc.GetRoutineByID(ctx, req.ID)
// 		case strings.TrimSpace(req.Name) != "":
// 			resp.Routine, resp.Err = svc.GetRoutineByName(ctx, req.Name)
// 		}

// 		return resp, nil
// 	}
// }

// func makeSearchRoutineEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(searchRoutinesByTagsRequest)
// 		resp := searchRoutinesByTagsResponse{}

// 		resp.Routines, resp.Err = svc.GetRoutinesByTags(ctx, req.Tags...)

// 		return resp, nil
// 	}
// }

// type addRoutineRequest struct {
// 	Name        string          `json:"name,omitempty"`
// 	Description string          `json:"description,omitempty"`
// 	Tags        []*internal.Tag `json:"tags,omitempty"`
// }

// type addRoutineResponse struct {
// 	Err error
// 	ID  int
// }

// type putRoutineRequest struct {
// 	ID      int
// 	Routine internal.Routine
// }

// type deleteRoutineRequest struct {
// 	ID   int    `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// }

// type getRoutineRequest struct {
// 	ID   int    `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// }

// type getRoutineResponse struct {
// 	Err     error             `json:"err,omitempty"`
// 	Routine *internal.Routine `json:"Routines,omitempty"`
// }

// type searchRoutinesByTagsRequest struct {
// 	Tags []*internal.Tag `json:"tags,omitempty"`
// }

// type searchRoutinesByTagsResponse struct {
// 	Err      error               `json:"err,omitempty"`
// 	Routines []*internal.Routine `json:"Routines,omitempty"`
// }

// type defaultResponse struct {
// 	Err error `json:"err,omitempty"`
// }
