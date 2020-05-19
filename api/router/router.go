package router

import (
	"github.com/davila23/jwt-auth/api/router/routes"
	"github.com/gorilla/mux"
)

// New routes
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
