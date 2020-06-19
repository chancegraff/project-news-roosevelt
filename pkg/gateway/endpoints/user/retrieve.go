package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	service "github.com/pronuu/roosevelt/pkg/gateway/services/user"
	transport "github.com/pronuu/roosevelt/pkg/gateway/transports/user"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.RetrieveRequest)
		user, err := svc.Retrieve(req.ID, req.Email)
		if err != nil {
			return transport.RetrieveResponse{
				User: *user,
				Err:  err.Error(),
			}, nil
		}
		return transport.RetrieveResponse{
			User: *user,
			Err:  "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, id uint, email string) (*models.User, error) {
	req := transport.RetrieveRequest{ID: id, Email: email}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(transport.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return &retrieveResp.User, nil
}
