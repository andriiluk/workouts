package workoutsvc

// import (
// 	"context"
// 	"strings"

// 	"github.com/andriiluk/workouts/internal"
// 	"github.com/go-kit/kit/endpoint"
// )

// type Endpoints struct {
// 	PostWorkout    endpoint.Endpoint
// 	PutWorkout     endpoint.Endpoint
// 	DeleteWorkout  endpoint.Endpoint
// 	GetWorkout     endpoint.Endpoint
// 	SearchWorkouts endpoint.Endpoint
// }

// func MakeEndpoints(svc Service) *Endpoints {
// 	return &Endpoints{
// 		PostWorkout:    makePostWorkoutEndpoint(svc),
// 		PutWorkout:     makePutWorkoutEndpoint(svc),
// 		DeleteWorkout:  makeDeleteWorkoutEndpoint(svc),
// 		GetWorkout:     makeGetWorkoutEndpoint(svc),
// 		SearchWorkouts: makeSearchWorkoutEndpoint(svc),
// 	}
// }

// func makePostWorkoutEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		var resp PostWorkoutResponse

// 		req := request.(PostWorkoutRequest)

// 		id, err := svc.AddWorkout(ctx, &internal.Workout{
// 			Name:        req.Name,
// 			Description: req.Description,
// 			Tags:        req.Tags,
// 		})
// 		if err != nil {
// 			resp.Err = err.Error()
// 		}

// 		return PostWorkoutResponse{
// 			ID: id,
// 		}, nil
// 	}
// }

// func makePutWorkoutEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(PutWorkoutRequest)

// 		err = svc.PutWorkout(ctx, &internal.Workout{
// 			ID:          req.ID,
// 			Name:        req.Workout.Name,
// 			Description: req.Workout.Description,
// 			Tags:        req.Workout.Tags,
// 		})

// 		return defaultResponse{
// 			Err: err.Error(),
// 		}, nil
// 	}
// }

// func makeDeleteWorkoutEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(DeleteWorkoutRequest)
// 		resp := defaultResponse{}

// 		var err error

// 		switch {
// 		case req.ID != 0:
// 			err = svc.DeleteWorkoutByID(ctx, req.ID)
// 		case strings.TrimSpace(req.Name) != "":
// 			err = svc.DeleteWorkoutByName(ctx, req.Name)
// 		}

// 		if err != nil {
// 			resp.Err = err.Error()
// 		}

// 		return resp, nil
// 	}
// }

// func makeGetWorkoutEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(getWorkoutRequest)
// 		resp := getWorkoutResponse{}

// 		switch {
// 		case req.ID != 0:
// 			resp.Workout, resp.Err = svc.GetWorkoutByID(ctx, req.ID)
// 		case strings.TrimSpace(req.Name) != "":
// 			resp.Workout, resp.Err = svc.GetWorkoutByName(ctx, req.Name)
// 		}

// 		return resp, nil
// 	}
// }

// func makeSearchWorkoutEndpoint(svc Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(searchWorkoutsByTagsRequest)
// 		resp := searchWorkoutsByTagsResponse{}

// 		resp.Workouts, resp.Err = svc.GetWorkoutsByTags(ctx, req.Tags...)

// 		return resp, nil
// 	}
// }

// type PostWorkoutRequest struct {
// 	Name        string          `json:"name,omitempty"`
// 	Description string          `json:"description,omitempty"`
// 	Tags        []*internal.Tag `json:"tags,omitempty"`
// }

// type PostWorkoutResponse struct {
// 	Err string `json:"err,omitempty"`
// 	ID  int    `json:"id,omitempty"`
// }

// type PutWorkoutRequest struct {
// 	Workout internal.Workout
// 	ID      int
// }

// type DeleteWorkoutRequest struct {
// 	Name string `json:"name,omitempty"`
// 	ID   int    `json:"id,omitempty"`
// }

// type getWorkoutRequest struct {
// 	Name string `json:"name,omitempty"`
// 	ID   int    `json:"id,omitempty"`
// }

// type getWorkoutResponse struct {
// 	Err     error             `json:"err,omitempty"`
// 	Workout *internal.Workout `json:"workouts,omitempty"`
// }

// type searchWorkoutsByTagsRequest struct {
// 	Tags []*internal.Tag `json:"tags,omitempty"`
// }

// type searchWorkoutsByTagsResponse struct {
// 	Err      error               `json:"err,omitempty"`
// 	Workouts []*internal.Workout `json:"workouts,omitempty"`
// }

// type defaultResponse struct {
// 	Err string `json:"err,omitempty"`
// }
