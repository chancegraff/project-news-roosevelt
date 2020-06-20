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
	Distinctions models.Distinctions `json:"client"`
}

// CreateResponse ...
type CreateResponse struct {
	Client models.Client `json:"client"`
	Err    string        `json:"err,omitempty"`
}

// DecodeCreateHTTPRequest ...
func DecodeCreateHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeCreateHTTPRequest ...
func EncodeCreateHTTPRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
