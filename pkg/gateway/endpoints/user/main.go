package user

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/pkg/gateway/services/user"
)

// Endpoints ...
type Endpoints struct {
	CreateEndpoint   endpoint.Endpoint
	UpdateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s user.Service) Endpoints {
	return Endpoints{
		CreateEndpoint:   MakeCreateEndpoint(s),
		UpdateEndpoint:   MakeUpdateEndpoint(s),
		RetrieveEndpoint: MakeRetrieveEndpoint(s),
	}
}
