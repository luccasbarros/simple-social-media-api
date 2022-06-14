package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate router with all configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
