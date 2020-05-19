package routes

import (
	"net/http"

	"github.com/davila23/jwt-auth/api/controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
