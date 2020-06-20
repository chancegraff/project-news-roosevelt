package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/pkg/user/services"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transports.CreateRequest)
		user, err := svc.Create(ctx, req.Credentials)
		if err != nil {
			return transports.CreateResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return transports.CreateResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, credentials models.Credentials) (*models.User, error) {
	req := transports.CreateRequest{Credentials: credentials}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(transports.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return createResp.User, nil
}
