package client

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/roosevelt/internal/models"
	service "github.com/pronuu/roosevelt/pkg/gateway/services/client"
	transport "github.com/pronuu/roosevelt/pkg/gateway/transports/client"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.RetrieveRequest)
		client, err := svc.Retrieve(req.ID, req.IP, req.Hash, req.UserID)
		if err != nil {
			return transport.RetrieveResponse{
				Client: *client,
				Err:    err.Error(),
			}, nil
		}
		return transport.RetrieveResponse{
			Client: *client,
			Err:    "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, id uint, ip, hash, userID string) (*models.Client, error) {
	req := transport.RetrieveRequest{ID: id, IP: ip, Hash: hash, UserID: userID}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(transport.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return &retrieveResp.Client, nil
}
