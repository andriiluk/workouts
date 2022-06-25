package exercisesvc

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

// 	// POST    /exercises                          	adds another profile
// 	// PUT     /exercises/:id                       post updated profile information about the profile
// 	// GET     /exercises/:id                       retrieves the given profile by id
// 	// DELETE  /exercises/:id                       remove the given profile
// 	// GET     /exercises/:id/addresses/            retrieve addresses associated with the profile
// 	// GET     /exercises/:id/addresses/:addressID  retrieve a particular profile address
// 	// POST    /exercises/:id/addresses/            add a new address
// 	// DELETE  /exercises/:id/addresses/:addressID  remove an address

// 	r.Methods("POST").Path("/exercises/").Handler(httptransport.NewServer(
// 		e.AddExercise,
// 		decodePostExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("PUT").Path("/exercises/:id").Handler(httptransport.NewServer(
// 		e.PutExercise,
// 		decodePutExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("DELETE").Path("/exercises/:id").Handler(httptransport.NewServer(
// 		e.DeleteExercise,
// 		decodeDeleteExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("GET").Path("/exercises/:id").Handler(httptransport.NewServer(
// 		e.GetExercise,
// 		decodeGetExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("GET").Path("/exercises/name/:name").Handler(httptransport.NewServer(
// 		e.GetExercise,
// 		decodeGetByNameExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	r.Methods("POST").Path("/exercises/search").Handler(httptransport.NewServer(
// 		e.GetExercise,
// 		decodeSearchExerciseRequest,
// 		encodeResponse,
// 		options...,
// 	))

// 	return r
// }

// func decodePostExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	var req addExerciseRequest
// 	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
// 		return nil, e
// 	}
// 	return req, nil
// }

// func decodePutExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	var exercise internal.Exercise

// 	if e := json.NewDecoder(r.Body).Decode(&exercise); e != nil {
// 		return nil, e
// 	}

// 	return putExerciseRequest{
// 		ID:       id,
// 		Exercise: exercise,
// 	}, nil
// }

// func decodeDeleteExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	return deleteExerciseRequest{
// 		ID: id,
// 	}, nil
// }

// func decodeGetExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	idStr, ok := mux.Vars(r)["id"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
// 	}

// 	return getExerciseRequest{
// 		ID: id,
// 	}, nil
// }

// func decodeGetByNameExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
// 	name, ok := mux.Vars(r)["name"]
// 	if !ok {
// 		return nil, ErrBadRequest
// 	}

// 	return getExerciseRequest{
// 		Name: name,
// 	}, nil
// }

// func decodeSearchExerciseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req searchExercisesByTagsRequest

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
