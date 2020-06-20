package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pronuu/roosevelt/pkg/client/endpoints"
)

// Routes ...
type Routes struct {
	CreateRoute   *httptransport.Server
	RetrieveRoute *httptransport.Server
}

// NewRoutes ...
func NewRoutes(endpoints *endpoints.Endpoints) Routes {
	return Routes{
		CreateRoute:   MakeCreateRoute(endpoints),
		RetrieveRoute: MakeRetrieveRoute(endpoints),
	}
}

// Route ...
func (r *Routes) Route(mxr *mux.Router) *mux.Router {
	route := mxr.PathPrefix("/client").Subrouter()
	route.HandleFunc("/create", r.Create).Methods("POST", "OPTIONS")
	route.HandleFunc("/retrieve", r.Retrieve).Methods("POST", "OPTIONS")
	return mxr
}
