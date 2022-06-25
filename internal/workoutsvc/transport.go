package workoutsvc

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/andriiluk/workouts/internal"
// 	"github.com/gorilla/mux"

// 	"github.com/go-kit/log"

// 	"github.com/go-kit/kit/transport"

// 	httptransport "github.com/go-kit/kit/transport/http"
// )

// var (
// 	ErrBadRequest = errors.New("bad request")
// )

// func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
// 	r := mux.NewRouter()
// 	e := MakeEndpoints(s)
// 	options := []httptransport.ServerOption{
// 		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
// 		httptransport.ServerErrorEncoder(encodeError),
// 	}

// 	// POST    /workouts                          	adds another profile
// 	// PUT     /workouts/:id                       post updated profile information about the profile
// 	// GET     /workouts/:id                       retrieves the given profile by id
// 	// DELETE  /workouts/:id                       remove the given profile
// 	// GET     /workouts/:id/addresses/            retrieve addresses associated with the profile
// 	// GET     /workouts/:id/addresses/:addressID  retrieve a particular profile address
// 	// POST    /workouts/:id/addresses/            add a new address
// 	// DELETE  /workouts/:id/addresses/:addressID  remove an address

// 	r.Methods("POST").Path("/workouts/").Handler(httptransport.NewServer(
// 		e.PostWorkout,
// 		decodePostWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("PUT").Path("/workouts/:id").Handler(httptransport.NewServer(
// 		e.PutWorkout,
// 		decodePutWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("DELETE").Path("/workouts/:id").Handler(httptransport.NewServer(
// 		e.DeleteWorkout,
// 		decodeDeleteWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("GET").Path("/workouts/:id").Handler(httptransport.NewServer(
// 		e.GetWorkout,
// 		decodeGetWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("GET").Path("/workouts/name/:name").Handler(httptransport.NewServer(
// 		e.GetWorkout,
// 		decodeGetByNameWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("POST").Path("/workouts/search").Handler(httptransport.NewServer(
// 		e.GetWorkout,
// 		decodeSearchWorkoutRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	return r
// }

// func decodePostWorkoutRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	var req PostWorkoutRequest
// 	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
// 		return nil, e
// 	}
// 	return req, nil
// }

// func decodePutWorkoutRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	var Workout internal.Workout

// 	if e := json.NewDecoder(r.Body).Decode(&Workout); e != nil {
// 		return nil, e
// 	}

// 	return PutWorkoutRequest{
// 		ID:      id,
// 		Workout: Workout,
// 	}, nil
// }

// func decodeDeleteWorkoutRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	return DeleteWorkoutRequest{
// 		ID: id,
// 	}, nil
// }

// func decodeGetWorkoutRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	return getWorkoutRequest{
// 		ID: id,
// 	}, nil
// }

// func decodeGetByNameWorkoutRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	name, ok := mux.Vars(r)["name"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	return getWorkoutRequest{
// 		Name: name,
// 	}, nil
// }

// func decodeSearchWorkoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req searchWorkoutsByTagsRequest

// 	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
// 		return nil, ErrBadRequest
// 	}

// 	return req, nil
// }

// func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")

// 	return json.NewEncoder(w).Encode(response)
// }

// func encodeError(_ context.Context, err error, w http.ResponseWriter) {
// 	if err == nil {
// 		panic("encodeError with nil error")
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(codeFrom(err))

// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"error": err.Error(),
// 	})
// }

// func codeFrom(err error) int {
// 	unErr := errors.Unwrap(err)

// 	switch {
// 	case errors.Is(ErrBadRequest, unErr):
// 		return http.StatusBadRequest
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
