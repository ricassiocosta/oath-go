package routes

import (
	"net/http"
	"oath-go/src/controllers"
)

var certificateRoutes = []Route{
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
		RequireAuth: false,
	},
	{
		URI:         "/signcertificate",
		Method:      http.MethodPost,
		Function:    controllers.SignCertificate,
		RequireAuth: false,
	},
}
