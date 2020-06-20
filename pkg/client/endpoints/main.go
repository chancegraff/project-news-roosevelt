package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/pkg/client/services"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	CreateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s services.Services) Endpoints {
	return Endpoints{
		CreateEndpoint:   MakeCreateEndpoint(s),
		RetrieveEndpoint: MakeRetrieveEndpoint(s),
	}
}

// Middleware ...
type Middleware func(Endpoints) Endpoints
