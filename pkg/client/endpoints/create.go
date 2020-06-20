package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/pkg/client/services"
	"github.com/pronuu/roosevelt/pkg/client/transports"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.CreateRequest)
		client, err := svc.Create(ctx, req.Distinctions)
		if err != nil {
			return transports.CreateResponse{
				Client: client,
				Err:    err.Error(),
			}, nil
		}
		return transports.CreateResponse{
			Client: client,
			Err:    "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, distinctions models.Distinctions) (*models.Client, error) {
	req := transports.CreateRequest{Distinctions: distinctions}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(transports.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return createResp.Client, nil
}
