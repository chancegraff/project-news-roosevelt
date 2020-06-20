package client

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/pkg/client/endpoints"
	"github.com/pronuu/roosevelt/pkg/client/middlewares"
	routes "github.com/pronuu/roosevelt/pkg/client/routes/http"
	"github.com/pronuu/roosevelt/pkg/client/services"
)

// NewClientService ...
func NewClientService(str *database.Store, lgr log.Logger) routes.Routes {
	mdl := middlewares.NewMiddleware(lgr)
	svc := services.NewServices(str)
	svc = mdl.BindServices(svc)
	end := endpoints.NewEndpoints(svc)
	end = mdl.BindEndpoints(end)
	rts := routes.NewRoutes(&end)
	return rts
}
