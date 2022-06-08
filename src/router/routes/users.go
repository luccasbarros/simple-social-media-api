package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
