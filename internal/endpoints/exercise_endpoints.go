package endpoints

import "github.com/go-kit/kit/endpoint"

type Exercise struct {
	AddEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}
