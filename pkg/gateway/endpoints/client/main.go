package client

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/pkg/gateway/services/client"
)

// Endpoints implements the endpoints for the service
type Endpoints struct {
	CreateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
}

// NewEndpoints instantiates the endpoints for the service
func NewEndpoints(s client.Service) Endpoints {
	return Endpoints{
		CreateEndpoint:   MakeCreateEndpoint(s),
		RetrieveEndpoint: MakeRetrieveEndpoint(s),
	}
}
