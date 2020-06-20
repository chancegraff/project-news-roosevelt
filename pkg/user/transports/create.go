package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/roosevelt/internal/models"
)

// CreateRequest ...
type CreateRequest struct {
	Credentials models.Credentials `json:"user"`
}

// CreateResponse ...
type CreateResponse struct {
	User *models.User `json:"user"`
	Err  string       `json:"err,omitempty"`
}

// DecodeCreateHTTPRequest ...
func DecodeCreateHTTPRequest(ctx context.Context, rq *http.Request) (interface{}, error) {
	var request CreateRequest
	if err := json.NewDecoder(rq.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeCreateHTTPResponse ...
func EncodeCreateHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
