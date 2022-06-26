package exercisesvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andriiluk/workouts/internal"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	exercisesPath = "/exercises"
)

// MakeHTTPHandler - creates http.Handler with below actions
// POST 	/exercises
// PUT 		/exercises/:id
// DELETE 	/exercises/:id
// GET 		/exercises/:id
// GET 		/exercises/name/:name
// POST		/exercises/search
func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(ErrorLogHandler),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path(exercisesPath).Handler(httptransport.NewServer(
		e.PostExerciseEndpoint,
		decodePostExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path(exercisesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.GetExerciseEndpoint,
		decodeGetExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path(exercisesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.DeleteExerciseEndpoint,
		decodeDeleteExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path(exercisesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.PutExerciseEndpoint,
		decodePutExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path(exercisesPath + "/name/:name").Handler(httptransport.NewServer(
		e.GetExerciseEndpoint,
		decodeGetByNameExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path(exercisesPath + "/search").Handler(httptransport.NewServer(
		e.SearchExercisesEndpoint,
		decodeSearchExerciseRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path(exercisesPath + "/bymuscles").Handler(httptransport.NewServer(
		e.GetExercisesByMusclesEndpoint,
		decodeGetExercisesByMusclesRequest,
		encodeResponse,
		options...,
	))

	return r
}

// ENCODERS
// ===================================

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	logrus.WithField("response", response).Debug("encode response")

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	if e := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	}); e != nil {
		panic(e)
	}
}

// DECODERS
// ===================================

func decodePostExerciseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PostExerciseRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		logrus.WithError(e).Debug()

		return nil, fmt.Errorf("[%v]: [%w]", internal.ErrDecodeRequest, e)
	}

	return req, nil
}

func decodePutExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, internal.ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", internal.ErrBadRequest, idStr)
	}

	var m *internal.Exercise

	if e := json.NewDecoder(r.Body).Decode(&m); e != nil {
		return nil, e
	}

	return PutExerciseRequest{
		ID:       id,
		Exercise: m,
	}, nil
}

func decodeDeleteExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, internal.ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", internal.ErrBadRequest, idStr)
	}

	return DeleteExerciseRequest{
		ID: id,
	}, nil
}

func decodeGetExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, internal.ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", internal.ErrBadRequest, idStr)
	}

	return GetExerciseRequest{
		ID: id,
	}, nil
}

func decodeGetByNameExerciseRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		return nil, internal.ErrBadRequest
	}

	return GetExerciseRequest{
		Name: name,
	}, nil
}

func decodeSearchExerciseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req SearchExercisesByTagsRequest

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, internal.ErrBadRequest
	}

	return req, nil
}

func decodeGetExercisesByMusclesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetExercisesByMusclesRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, internal.ErrBadRequest
	}

	return req, nil
}

func codeFrom(err error) int {
	unErr := errors.Unwrap(err)

	switch {
	case errors.Is(internal.ErrBadRequest, unErr):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
