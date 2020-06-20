package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	"github.com/pronuu/roosevelt/pkg/user/services"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.UpdateRequest)
		user, err := svc.Update(ctx, req.ID, req.Credentials)
		if err != nil {
			return &transports.UpdateResponse{
				User: user,
				Err:  err.Error(),
			}, nil
		}
		return &transports.UpdateResponse{
			User: user,
			Err:  "",
		}, nil
	}
}

// Update ...
func (e *Endpoints) Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error) {
	req := &transports.UpdateRequest{ID: id, Credentials: credentials}
	resp, err := e.UpdateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	updateResp := resp.(*transports.UpdateResponse)
	if updateResp.Err != "" {
		return nil, errors.New(updateResp.Err)
	}
	return updateResp.User, nil
}
