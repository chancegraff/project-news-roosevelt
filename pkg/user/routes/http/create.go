package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/roosevelt/internal/utils"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeCreateRoute ...
func MakeCreateRoute(endpoints endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.CreateEndpoint,
		transports.DecodeCreateHTTPRequest,
		transports.EncodeCreateHTTPResponse,
	)
}

// Create ...
func (r *Routes) Create(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.CreateRoute.ServeHTTP(rw, rq)
}
