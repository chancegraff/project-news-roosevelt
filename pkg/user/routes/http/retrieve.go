package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/roosevelt/internal/utils"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	"github.com/pronuu/roosevelt/pkg/user/transports"
)

// MakeRetrieveRoute ...
func MakeRetrieveRoute(endpoints *endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.RetrieveEndpoint,
		transports.DecodeRetrieveHTTPRequest,
		transports.EncodeRetrieveHTTPRequest,
	)
}

// Retrieve ...
func (r *Routes) Retrieve(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.RetrieveRoute.ServeHTTP(rw, rq)
}
