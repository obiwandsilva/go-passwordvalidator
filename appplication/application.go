package appplication

import (
	"github.com/obiwandsilva/passwordvalidator/appplication/config"
	"github.com/obiwandsilva/passwordvalidator/appplication/web"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/controllers"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/routes"
)

type PasswordValidator struct {
	EnvConfig config.EnvironmentConfig
	PasswordValidatorController controllers.PasswordValidatorController
}

func (pw *PasswordValidator) Start() {
	server := web.Server{
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	routes := routes.Routes(pw.PasswordValidatorController)

	server.Start("8080", routes)
}
