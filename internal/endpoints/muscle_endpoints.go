package endpoints

import (
	"github.com/go-kit/kit/endpoint"
)

type Muscle struct {
	AddEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}
