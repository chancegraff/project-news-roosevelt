package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	service "github.com/pronuu/roosevelt/pkg/gateway/services/user"
	transport "github.com/pronuu/roosevelt/pkg/gateway/transports/user"
)

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.UpdateRequest)
		user, err := svc.Update(req.ID, req.Credentials)
		if err != nil {
			return transport.UpdateResponse{
				User: *user,
				Err:  err.Error(),
			}, nil
		}
		return transport.UpdateResponse{
			User: *user,
			Err:  "",
		}, nil
	}
}

// Update ...
func (e *Endpoints) Update(ctx context.Context, id uint, credentials models.Credentials) (*models.User, error) {
	req := transport.UpdateRequest{ID: id, Credentials: credentials}
	resp, err := e.UpdateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	updateResp := resp.(transport.UpdateResponse)
	if updateResp.Err != "" {
		return nil, errors.New(updateResp.Err)
	}
	return &updateResp.User, nil
}
