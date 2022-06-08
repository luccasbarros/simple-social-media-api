package routes

import "net/http"

// All API Routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}
