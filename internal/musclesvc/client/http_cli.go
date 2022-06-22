package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/sirupsen/logrus"

	"github.com/andriiluk/workouts/internal"
	"github.com/andriiluk/workouts/internal/musclesvc"
)

func MakeClientEndpoints(svcAddr string) (*musclesvc.Endpoints, error) {
	if !strings.HasPrefix(svcAddr, "http") {
		svcAddr = "http://" + svcAddr
	}
	svcAddr += "/muscles"
	tgt, err := url.Parse(svcAddr)
	if err != nil {
		return nil, err
	}

	options := []httptransport.ClientOption{}
	log.WithFields(log.Fields{"url": tgt}).Debug("make client endpoints")

	// cliURL, _ := url.Parse("http://localhost:8000/muscles")

	return &musclesvc.Endpoints{
		PostMuscleEndpoint:    httptransport.NewClient("POST", copyURL(tgt), encodeRequest, decodePostMuscleResponse, options...).Endpoint(),
		PutMuscleEndpoint:     httptransport.NewClient(http.MethodPut, copyURL(tgt), encodePutMuscleRequest, decodeDefaultResponse, options...).Endpoint(),
		DeleteMuscleEndpoint:  httptransport.NewClient("DELETE", copyURL(tgt), encodeDeleteMuscleRequest, decodeDefaultResponse, options...).Endpoint(),
		GetMuscleEndpoint:     httptransport.NewClient("GET", copyURL(tgt), encodeGetMuscleRequest, decodeGetMuscleResponse, options...).Endpoint(),
		SearchMusclesEndpoint: httptransport.NewClient("POST", copyURL(tgt), encodeSearchMuscleRequest, decodeSearchMuscleResponse, options...).Endpoint(),
	}, nil
}

func copyURL(u *url.URL) *url.URL {
	newURL := &url.URL{}
	*newURL = *u

	return newURL
}

type HTTPClient struct {
	*musclesvc.Endpoints
}

func NewHTTPClient(svcAddr string) (*HTTPClient, error) {
	endpoints, err := MakeClientEndpoints(svcAddr)
	if err != nil {
		return nil, err
	}

	return &HTTPClient{endpoints}, nil
}

func (cli *HTTPClient) PostMuscle(name, description string, tags ...string) (int, error) {
	req := musclesvc.PostMuscleRequest{
		Name:        name,
		Description: description,
		Tags:        tags,
	}
	log.WithFields(log.Fields{"request": req}).Debug("post muscle request")

	resp, err := cli.PostMuscleEndpoint(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("endpoint error: [%w]", err)
	}

	postResp := resp.(musclesvc.PostMuscleResponse)
	return postResp.ID, nil
}

func (cli *HTTPClient) GetMuscle(id int) (*internal.Muscle, error) {
	req := musclesvc.GetMuscleRequest{
		ID: id,
	}
	resp, err := cli.GetMuscleEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}

	getResp := resp.(musclesvc.GetMuscleResponse)
	return getResp.Muscle, getResp.Err
}

func (cli *HTTPClient) DeleteMuscle(id int) error {
	req := musclesvc.DeleteMuscleRequest{
		ID: id,
	}
	resp, err := cli.DeleteMuscleEndpoint(context.Background(), req)
	if err != nil {
		return err
	}

	postResp := resp.(musclesvc.DefaultResponse)
	return postResp.Err
}

func (cli *HTTPClient) PutMuscle(m *internal.Muscle) error {
	req := musclesvc.PutMuscleRequest{
		ID:     m.ID,
		Muscle: m,
	}
	log.WithFields(log.Fields{"request": req}).Debug("cli put muscle")
	resp, err := cli.PutMuscleEndpoint(context.Background(), req)
	if err != nil {
		return err
	}

	putResp := resp.(musclesvc.DefaultResponse)
	return putResp.Err
}

func (cli *HTTPClient) SearchMusclesByTags(tags ...string) ([]*internal.Muscle, error) {
	if len(tags) == 0 {
		return nil, errors.New("empty tags")
	}

	req := musclesvc.SearchMusclesByTagsRequest{Tags: tags}
	resp, err := cli.SearchMusclesEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}

	searchResp := resp.(musclesvc.SearchMusclesByTagsResponse)
	return searchResp.Muscles, searchResp.Err
}

func encodeDeleteMuscleRequest(ctx context.Context, r *http.Request, request interface{}) error {
	delReq := request.(musclesvc.DeleteMuscleRequest)

	r.URL.Path += "/" + strconv.Itoa(delReq.ID)

	return encodeRequest(ctx, r, request)
}

func encodePutMuscleRequest(ctx context.Context, r *http.Request, request interface{}) error {
	putReq := request.(musclesvc.PutMuscleRequest)

	r.URL.Path += "/" + strconv.Itoa(putReq.ID)
	log.WithFields(log.Fields{"url": r.URL, "request": putReq}).Debug("encode put muscle request")
	return encodeRequest(ctx, r, putReq.Muscle)
}

func encodeGetMuscleRequest(ctx context.Context, r *http.Request, request interface{}) error {
	getReq := request.(musclesvc.GetMuscleRequest)

	r.URL.Path += "/" + strconv.Itoa(getReq.ID)
	log.WithFields(log.Fields{"url": r.URL, "request": getReq}).Debug("encode get muscle request")

	return encodeRequest(ctx, r, getReq)
}

func encodeSearchMuscleRequest(ctx context.Context, r *http.Request, request interface{}) error {
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

func decodePostMuscleResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var addMuscleResp musclesvc.PostMuscleResponse

	if e := json.NewDecoder(resp.Body).Decode(&addMuscleResp); e != nil {
		return nil, e
	}

	return addMuscleResp, nil
}

func decodeGetMuscleResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var getMuscleResp musclesvc.GetMuscleResponse

	if e := json.NewDecoder(resp.Body).Decode(&getMuscleResp); e != nil {
		return nil, e
	}

	return getMuscleResp, nil
}

func decodeSearchMuscleResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var searchMuscleResp musclesvc.SearchMusclesByTagsResponse

	if e := json.NewDecoder(resp.Body).Decode(&searchMuscleResp); e != nil {
		return nil, e
	}

	return searchMuscleResp, nil
}

func decodeDefaultResponse(ctx context.Context, resp *http.Response) (response interface{}, err error) {
	var defResponse musclesvc.DefaultResponse

	if e := json.NewDecoder(resp.Body).Decode(&defResponse); e != nil {
		return nil, e
	}

	return defResponse, nil
}
