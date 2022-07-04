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
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowers,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/users/{id}/following",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowing,
		AuthenticationRequired: true,
	},
}
