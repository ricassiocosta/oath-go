package routes

import (
	"net/http"
	"oath-go/src/controllers"
)

var authRoutes = []Route{
	{
		URI:         "/callback/github",
		Method:      http.MethodGet,
		Function:    controllers.GithubCallback,
		RequireAuth: false,
	},
	{
		URI:         "/auth",
		Method:      http.MethodGet,
		Function:    controllers.Auth,
		RequireAuth: false,
	},
}
