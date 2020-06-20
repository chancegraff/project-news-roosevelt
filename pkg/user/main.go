package user

import (
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
	routes "github.com/pronuu/roosevelt/pkg/user/routes/http"
	"github.com/pronuu/roosevelt/pkg/user/services"
)

// NewUserService ...
func NewUserService(str *database.Store) routes.Routes {
	svc := services.NewService(str)
	// svc = mdl.BindService(svc)
	end := endpoints.NewEndpoints(svc)
	// end = mdl.BindEndpoints(end)
	rts := routes.NewRoutes(&end)
	return rts
}
