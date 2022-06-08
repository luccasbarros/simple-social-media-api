package router

import "github.com/gorilla/mux"

// Generate router with all configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
