package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/roosevelt/internal/models"
)

// RetrieveRequest ...
type RetrieveRequest struct {
	User *models.User `json:"user"`
}

// RetrieveResponse ...
type RetrieveResponse struct {
	User *models.User `json:"user"`
	Err  string       `json:"err,omitempty"`
}

// DecodeRetrieveHTTPRequest ...
func DecodeRetrieveHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RetrieveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeRetrieveHTTPResponse ...
func EncodeRetrieveHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
