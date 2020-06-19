package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	service "github.com/pronuu/roosevelt/pkg/gateway/services/client"
	transport "github.com/pronuu/roosevelt/pkg/gateway/transports/client"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CreateRequest)
		client, err := svc.Create(req.Distinctions)
		if err != nil {
			return transport.CreateResponse{
				Client: *client,
				Err:    err.Error(),
			}, nil
		}
		return transport.CreateResponse{
			Client: *client,
			Err:    "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, distinctions models.Distinctions) (*models.Client, error) {
	req := transport.CreateRequest{Distinctions: distinctions}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(transport.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return &createResp.Client, nil
}
