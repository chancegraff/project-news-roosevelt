package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pronuu/roosevelt/pkg/user/endpoints"
)

// Routes ...
type Routes struct {
	CreateRoute   *httptransport.Server
	UpdateRoute   *httptransport.Server
	RetrieveRoute *httptransport.Server
}

// NewRoutes ...
func NewRoutes(endpoints *endpoints.Endpoints) Routes {
	return Routes{
		CreateRoute:   MakeCreateRoute(endpoints),
		UpdateRoute:   MakeUpdateRoute(endpoints),
		RetrieveRoute: MakeRetrieveRoute(endpoints),
	}
}

// Route ...
func (r *Routes) Route(mxr *mux.Router) *mux.Router {
	route := mxr.PathPrefix("/user").Subrouter()
	route.HandleFunc("/create", r.Create).Methods("POST", "OPTIONS")
	route.HandleFunc("/update", r.Update).Methods("POST", "OPTIONS")
	route.HandleFunc("/retrieve", r.Retrieve).Methods("POST", "OPTIONS")
	return mxr
}
