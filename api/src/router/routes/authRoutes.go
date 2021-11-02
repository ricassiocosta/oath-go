package routes

import (
	"net/http"
	"oath-go/src/controllers"
)

var authRoutes = []Route{
	{
		URI:         "/callback/github",
		Method:      http.MethodGet,
		Function:    controllers.AuthGithub,
		RequireAuth: false,
	},
}
