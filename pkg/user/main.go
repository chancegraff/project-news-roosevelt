package user

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	"github.com/pronuu/roosevelt/pkg/user/middlewares"
	routes "github.com/pronuu/roosevelt/pkg/user/routes/http"
	"github.com/pronuu/roosevelt/pkg/user/services"
)

// NewUserService ...
func NewUserService(str *database.Store, lgr log.Logger) routes.Routes {
	mdl := middlewares.NewMiddleware(lgr)
	svc := services.NewServices(str)
	svc = mdl.BindServices(svc)
	end := endpoints.NewEndpoints(svc)
	end = mdl.BindEndpoints(end)
	rts := routes.NewRoutes(end)
	return rts
}
