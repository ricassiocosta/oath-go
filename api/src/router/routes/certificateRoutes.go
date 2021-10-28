package routes

import (
	"net/http"
	"oath-go/src/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/crl",
		Method:      http.MethodGet,
		Function:    controllers.GetCRL,
		RequireAuth: false,
	},
	{
		URI:         "/crl",
		Method:      http.MethodPost,
		Function:    controllers.AddCertToCRL,
		RequireAuth: true,
	},
	{
		URI:         "/trustbundle",
		Method:      http.MethodGet,
		Function:    controllers.GetTrustBundle,
		RequireAuth: true,
	},
}
