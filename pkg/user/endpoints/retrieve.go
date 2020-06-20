package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/pkg/user/services"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.RetrieveRequest)
		user, err := svc.Retrieve(ctx, req.User)
		if err != nil {
			return &transports.RetrieveResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return &transports.RetrieveResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, user *models.User) (*models.User, error) {
	req := &transports.RetrieveRequest{User: user}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(*transports.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return retrieveResp.User, nil
}
