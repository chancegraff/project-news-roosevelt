package client

import (
	"github.com/pronuu/roosevelt/internal/database"
	"github.com/pronuu/roosevelt/pkg/client/endpoints"
	routes "github.com/pronuu/roosevelt/pkg/client/routes/http"
	"github.com/pronuu/roosevelt/pkg/client/services"
)

// NewClientService ...
func NewClientService(str *database.Store) routes.Routes {
	svc := services.NewService(str)
	// svc = mdl.BindService(svc)
	end := endpoints.NewEndpoints(svc)
	// end = mdl.BindEndpoints(end)
	rts := routes.NewRoutes(&end)
	return rts
}
