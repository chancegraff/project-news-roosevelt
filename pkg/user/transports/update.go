package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/roosevelt/internal/models"
)

// UpdateRequest ...
type UpdateRequest struct {
	ID          uint               `json:"id"`
	Credentials models.Credentials `json:"user"`
}

// UpdateResponse ...
type UpdateResponse struct {
	User *models.User `json:"user"`
	Err  string       `json:"err,omitempty"`
}

// DecodeUpdateHTTPRequest ...
func DecodeUpdateHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeUpdateHTTPRequest ...
func EncodeUpdateHTTPRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
