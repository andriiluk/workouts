package routinesvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andriiluk/workouts/internal"
	"github.com/gorilla/mux"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrBadRequest = errors.New("bad request")
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// POST    /routines                          	adds another profile
	// PUT     /routines/:id                       post updated profile information about the profile
	// GET     /routines/:id                       retrieves the given profile by id
	// DELETE  /routines/:id                       remove the given profile
	// GET     /routines/:id/addresses/            retrieve addresses associated with the profile
	// GET     /routines/:id/addresses/:addressID  retrieve a particular profile address
	// POST    /routines/:id/addresses/            add a new address
	// DELETE  /routines/:id/addresses/:addressID  remove an address

	r.Methods("POST").Path("/routines/").Handler(httptransport.NewServer(
		e.AddRoutine,
		decodePostRoutineRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/routines/:id").Handler(httptransport.NewServer(
		e.PutRoutine,
		decodePutRoutineRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/routines/:id").Handler(httptransport.NewServer(
		e.DeleteRoutine,
		decodeDeleteRoutineRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/routines/:id").Handler(httptransport.NewServer(
		e.GetRoutine,
		decodeGetRoutineRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/routines/name/:name").Handler(httptransport.NewServer(
		e.GetRoutine,
		decodeGetByNameRoutineRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/routines/search").Handler(httptransport.NewServer(
		e.GetRoutine,
		decodeSearchRoutineRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodePostRoutineRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req addRoutineRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutRoutineRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	var Routine internal.Routine

	if e := json.NewDecoder(r.Body).Decode(&Routine); e != nil {
		return nil, e
	}

	return putRoutineRequest{
		ID:      id,
		Routine: Routine,
	}, nil
}

func decodeDeleteRoutineRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	return deleteRoutineRequest{
		ID: id,
	}, nil
}

func decodeGetRoutineRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	return getRoutineRequest{
		ID: id,
	}, nil
}

func decodeGetByNameRoutineRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		return nil, ErrBadRequest
	}

	return getRoutineRequest{
		Name: name,
	}, nil
}

func decodeSearchRoutineRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req searchRoutinesByTagsRequest

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, ErrBadRequest
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	unErr := errors.Unwrap(err)

	switch {
	case errors.Is(ErrBadRequest, unErr):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
