package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/sirupsen/logrus"

	"github.com/andriiluk/workouts/internal"
	"github.com/andriiluk/workouts/internal/exercisesvc"
)

func MakeClientEndpoints(svcAddr string) (*exercisesvc.Endpoints, error) {
	if !strings.HasPrefix(svcAddr, "http") {
		svcAddr = "http://" + svcAddr
	}

	svcAddr += "/exercises"

	tgt, err := url.Parse(svcAddr)
	if err != nil {
		return nil, err
	}

	options := []httptransport.ClientOption{}

	log.WithFields(log.Fields{"url": tgt}).Debug("make client endpoints")

	return &exercisesvc.Endpoints{
		PostExerciseEndpoint: httptransport.NewClient(
			http.MethodPost, copyURL(tgt), encodeRequest, decodePostExerciseResponse, options...,
		).Endpoint(),
		PutExerciseEndpoint: httptransport.NewClient(
			http.MethodPut, copyURL(tgt), encodePutExerciseRequest, decodeDefaultResponse, options...,
		).Endpoint(),
		DeleteExerciseEndpoint: httptransport.NewClient(
			http.MethodDelete, copyURL(tgt), encodeDeleteExerciseRequest, decodeDefaultResponse, options...,
		).Endpoint(),
		GetExerciseEndpoint: httptransport.NewClient(
			http.MethodGet, copyURL(tgt), encodeGetExerciseRequest, decodeGetExerciseResponse, options...,
		).Endpoint(),
		SearchExercisesEndpoint: httptransport.NewClient(
			http.MethodPost, copyURL(tgt), encodeSearchExerciseRequest, decodeSearchExerciseResponse, options...,
		).Endpoint(),
	}, nil
}

func copyURL(u *url.URL) *url.URL {
	newURL := &url.URL{}
	*newURL = *u

	return newURL
}

type HTTPClient struct {
	*exercisesvc.Endpoints
}

func NewHTTPClient(svcAddr string) (*HTTPClient, error) {
	endpoints, err := MakeClientEndpoints(svcAddr)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{endpoints}, nil
}

func (cli *HTTPClient) PostMuscle(name, description string, tags ...string) (int, error) {
	req := exercisesvc.PostExerciseRequest{
		Name:        name,
		Description: description,
		Tags:        tags,
	}

	log.WithFields(log.Fields{"request": req}).Debug("post muscle request")

	resp, err := cli.PostExerciseEndpoint(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("endpoint error: [%w]", err)
	}

	postResp, ok := resp.(exercisesvc.PostExerciseResponse)
	if !ok {
		return 0, fmt.Errorf("[%w]: expencted PostExerciseResponse", internal.ErrBadResponse)
	}

	return postResp.ID, nil
}

func (cli *HTTPClient) GetMuscle(id int) (*internal.Exercise, error) {
	req := exercisesvc.GetExerciseRequest{
		ID: id,
	}

	resp, err := cli.GetExerciseEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}

	getResp, ok := resp.(exercisesvc.GetExerciseResponse)
	if !ok {
		return nil, fmt.Errorf("[%w]: response expected to be GetExerciseResponse", internal.ErrBadResponse)
	}

	return getResp.Exercise, fmt.Errorf("[%w]: %s", internal.ErrInternalService, getResp.Err)
}

func (cli *HTTPClient) DeleteMuscle(id int) error {
	req := exercisesvc.DeleteExerciseRequest{
		ID: id,
	}

	resp, err := cli.DeleteExerciseEndpoint(context.Background(), req)
	if err != nil {
		return err
	}

	postResp, ok := resp.(exercisesvc.DefaultResponse)
	if !ok {
		return fmt.Errorf("[%w]: response expected to be DefaultResponse", internal.ErrBadResponse)
	}

	return postResp.Err
}

func (cli *HTTPClient) PutMuscle(m *internal.Exercise) error {
	req := exercisesvc.PutExerciseRequest{
		ID:       m.ID,
		Exercise: m,
	}

	log.WithFields(log.Fields{"request": req}).Debug("cli put muscle")

	resp, err := cli.PutExerciseEndpoint(context.Background(), req)
	if err != nil {
		return err
	}

	putResp, ok := resp.(exercisesvc.DefaultResponse)
	if !ok {
		return fmt.Errorf("[%w]: response expected to be DefaultResponse", internal.ErrBadResponse)
	}

	return putResp.Err
}

func (cli *HTTPClient) SearchExercisesByTags(tags ...string) ([]*internal.Exercise, error) {
	if len(tags) == 0 {
		return nil, fmt.Errorf("[%w]: empty tags", internal.ErrBadRequest)
	}

	req := exercisesvc.SearchExercisesByTagsRequest{Tags: tags}

	resp, err := cli.SearchExercisesEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}

	searchResp, ok := resp.(exercisesvc.SearchExercisesByTagsResponse)
	if !ok {
		return nil, fmt.Errorf("[%w]: response exptected to be SearchMusclesByTagsResponse", internal.ErrBadResponse)
	}

	return searchResp.Exercises, searchResp.Err
}

func encodeDeleteExerciseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	delReq, ok := request.(exercisesvc.DeleteExerciseRequest)
	if !ok {
		return fmt.Errorf("[%w]: request expected to be DeleteExerciseRequest", internal.ErrBadRequest)
	}

	r.URL.Path += "/" + strconv.Itoa(delReq.ID)

	return encodeRequest(ctx, r, request)
}

func encodePutExerciseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	putReq, ok := request.(exercisesvc.PutExerciseRequest)
	if !ok {
		return fmt.Errorf("[%w]: request expected to be PutExerciseRequest", internal.ErrBadRequest)
	}

	r.URL.Path += "/" + strconv.Itoa(putReq.ID)

	log.WithFields(log.Fields{"url": r.URL, "request": putReq}).Debug("encode put muscle request")

	return encodeRequest(ctx, r, putReq.Exercise)
}

func encodeGetExerciseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	getReq, ok := request.(exercisesvc.GetExerciseRequest)
	if !ok {
		return fmt.Errorf("[%w]: request exptected to be GetExerciseRequest", internal.ErrBadResponse)
	}

	r.URL.Path += "/" + strconv.Itoa(getReq.ID)

	log.WithFields(log.Fields{"url": r.URL, "request": getReq}).Debug("encode get muscle request")

	return encodeRequest(ctx, r, getReq)
}

func encodeSearchExerciseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path += "/search"

	return encodeRequest(ctx, r, request)
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(&buf)

	return nil
}

func decodePostExerciseResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var addMuscleResp exercisesvc.PostExerciseResponse

	if e := json.NewDecoder(resp.Body).Decode(&addMuscleResp); e != nil {
		return nil, e
	}

	return addMuscleResp, nil
}

func decodeGetExerciseResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var getMuscleResp exercisesvc.GetExerciseResponse

	if e := json.NewDecoder(resp.Body).Decode(&getMuscleResp); e != nil {
		return nil, e
	}

	return getMuscleResp, nil
}

func decodeSearchExerciseResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var searchMuscleResp exercisesvc.SearchExercisesByTagsResponse

	if e := json.NewDecoder(resp.Body).Decode(&searchMuscleResp); e != nil {
		return nil, e
	}

	return searchMuscleResp, nil
}

func decodeDefaultResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var defResponse exercisesvc.DefaultResponse

	if e := json.NewDecoder(resp.Body).Decode(&defResponse); e != nil {
		return nil, e
	}

	return defResponse, nil
}
