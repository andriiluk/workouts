package musclesvc

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

var (
	ErrBadRequest    = errors.New("bad request")
	ErrDecodeRequest = errors.New("error decoding request")
)

const (
	musclesPath = "/muscles"
)

// MakeHTTPHandler - creates http.Handler with below actions
// POST 	/muscles
// PUT 		/muscles/:id
// DELETE 	/muscles/:id
// GET 		/muscles/:id
// GET 		/muscles/name/:name
// POST		/muscles/search
func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(ErrorLogHandler),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path(musclesPath).Handler(httptransport.NewServer(
		e.PostMuscleEndpoint,
		decodePostMuscleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path(musclesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.GetMuscleEndpoint,
		decodeGetMuscleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path(musclesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.DeleteMuscleEndpoint,
		decodeDeleteMuscleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path(musclesPath + "/{id:[0-9]+}").Handler(httptransport.NewServer(
		e.PutMuscleEndpoint,
		decodePutMuscleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path(musclesPath + "/name/:name").Handler(httptransport.NewServer(
		e.GetMuscleEndpoint,
		decodeGetByNameMuscleRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path(musclesPath + "/search").Handler(httptransport.NewServer(
		e.SearchMusclesEndpoint,
		decodeSearchMuscleRequest,
		encodeResponse,
		options...,
	))

	return r
}

// ENCODERS
// ===================================

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	logrus.WithField("Response", response).Debug("encode response")

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

// DECODERS
// ===================================

func decodePostMuscleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PostMuscleRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		logrus.WithError(e).Debug()

		return nil, fmt.Errorf("[%v]: [%w]", ErrDecodeRequest, e)
	}

	return req, nil
}

func decodePutMuscleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	var m *internal.Muscle

	if e := json.NewDecoder(r.Body).Decode(&m); e != nil {
		return nil, e
	}

	return PutMuscleRequest{
		ID:     id,
		Muscle: m,
	}, nil
}

func decodeDeleteMuscleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	return DeleteMuscleRequest{
		ID: id,
	}, nil
}

func decodeGetMuscleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("[%w]. Invalid id: %s", ErrBadRequest, idStr)
	}

	return GetMuscleRequest{
		ID: id,
	}, nil
}

func decodeGetByNameMuscleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		return nil, ErrBadRequest
	}

	return GetMuscleRequest{
		Name: name,
	}, nil
}

func decodeSearchMuscleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req SearchMusclesByTagsRequest

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, ErrBadRequest
	}

	return req, nil
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
