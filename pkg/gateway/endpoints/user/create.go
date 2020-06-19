package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	service "github.com/pronuu/roosevelt/pkg/gateway/services/user"
	transport "github.com/pronuu/roosevelt/pkg/gateway/transports/user"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CreateRequest)
		user, err := svc.Create(req.Credentials)
		if err != nil {
			return transport.CreateResponse{
				User: *user,
				Err:  err.Error(),
			}, nil
		}
		return transport.CreateResponse{
			User: *user,
			Err:  "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, credentials models.Credentials) (*models.User, error) {
	req := transport.CreateRequest{Credentials: credentials}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(transport.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return &createResp.User, nil
}
