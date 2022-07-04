package routes

import "net/http"

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPersonalPosts,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{post_id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{post_id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/posts/{post_id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		AuthenticationRequired: true,
	},
}
