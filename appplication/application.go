package appplication

import (
	"github.com/obiwandsilva/passwordvalidator/appplication/config"
	"github.com/obiwandsilva/passwordvalidator/appplication/web"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/controllers"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/routes"
	"time"
)

type PasswordValidator struct {
	EnvConfig config.EnvironmentConfig
	PasswordValidatorController *controllers.PasswordValidatorController
}

func (pw *PasswordValidator) Start() {
	server := web.Server{
		ReadTimeout:  time.Second * time.Duration(pw.EnvConfig.ServerReadTimeout),
		WriteTimeout: time.Second * time.Duration(pw.EnvConfig.ServerWriteTimeout),
	}

	routes := routes.Routes(pw.PasswordValidatorController)

	server.Start(pw.EnvConfig.ServerPort, routes)
}
