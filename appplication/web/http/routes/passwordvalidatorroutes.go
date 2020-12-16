package routes

import (
	"github.com/bnkamalesh/webgo/v4"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/controllers"
	"net/http"
)

func Routes(passwordValidatorController controllers.PasswordValidatorController) []*webgo.Route {
	return []*webgo.Route{
		{
			Name: "Validate password",
			Method: http.MethodPost,
			Pattern: "/validate",
			Handlers: []http.HandlerFunc{
				passwordValidatorController.ValidatePassword,
			},
		},
	}
}