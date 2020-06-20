package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/pkg/client/services"
	"github.com/pronuu/roosevelt/pkg/client/transports"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.RetrieveRequest)
		client, err := svc.Retrieve(ctx, req.Client)
		if err != nil {
			return &transports.RetrieveResponse{
				Client: client,
				Err:    err.Error(),
			}, nil
		}
		return &transports.RetrieveResponse{
			Client: client,
			Err:    "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, client *models.Client) (*models.Client, error) {
	req := &transports.RetrieveRequest{Client: client}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(*transports.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return retrieveResp.Client, nil
}
