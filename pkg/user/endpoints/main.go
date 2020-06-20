package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/pkg/user/services"
)

// Endpoints ...
type Endpoints struct {
	CreateEndpoint   endpoint.Endpoint
	UpdateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s services.Service) Endpoints {
	return Endpoints{
		CreateEndpoint:   MakeCreateEndpoint(s),
		UpdateEndpoint:   MakeUpdateEndpoint(s),
		RetrieveEndpoint: MakeRetrieveEndpoint(s),
	}
}
