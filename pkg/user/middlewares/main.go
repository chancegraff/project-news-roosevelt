package middlewares

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	"github.com/pronuu/roosevelt/pkg/user/middlewares/application"
	"github.com/pronuu/roosevelt/pkg/user/middlewares/transport"
	"github.com/pronuu/roosevelt/pkg/user/services"
)

// NewMiddleware ...
func NewMiddleware(logger log.Logger) *Middleware {
	return &Middleware{
		ApplicationLogger: application.MakeMiddleware(logger),
		TransportLogger:   transport.MakeMiddleware(logger),
	}
}

// Middleware ...
type Middleware struct {
	ApplicationLogger services.Middleware
	TransportLogger   endpoints.Middleware
}

// BindServices ...
func (mw *Middleware) BindServices(svc services.Services) services.Services {
	return mw.ApplicationLogger(svc)
}

// BindEndpoints ...
func (mw *Middleware) BindEndpoints(end endpoints.Endpoints) endpoints.Endpoints {
	return mw.TransportLogger(end)
}
