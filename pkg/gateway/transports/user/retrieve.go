package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/roosevelt/internal/models"
)

// RetrieveRequest ...
type RetrieveRequest struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

// RetrieveResponse ...
type RetrieveResponse struct {
	User models.User `json:"user"`
	Err  string      `json:"err,omitempty"`
}

// DecodeRetrieveHTTPRequest ...
func DecodeRetrieveHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RetrieveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeRetrieveHTTPRequest ...
func EncodeRetrieveHTTPRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
