package transport

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
)

// MakeMiddleware will create the middlewares
func MakeMiddleware(lgr log.Logger) endpoints.Middleware {
	lgr = log.With(lgr, "layer", "transport")
	return func(end endpoints.Endpoints) endpoints.Endpoints {
		return endpoints.Endpoints{
			CreateEndpoint:   MakeCreateMiddleware(lgr, end),
			RetrieveEndpoint: MakeRetrieveMiddleware(lgr, end),
			UpdateEndpoint:   MakeUpdateMiddleware(lgr, end),
		}
	}
}
