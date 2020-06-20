package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/roosevelt/internal/utils"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeUpdateRoute ...
func MakeUpdateRoute(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.UpdateEndpoint,
		transports.DecodeUpdateHTTPRequest,
		transports.EncodeUpdateHTTPResponse,
	)
}

// Update ...
func (r *Routes) Update(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.UpdateRoute.ServeHTTP(rw, rq)
}
