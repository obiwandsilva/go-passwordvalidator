package application

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/obiwandsilva/passwordvalidator/application/config"
	"github.com/obiwandsilva/passwordvalidator/application/web"
	"github.com/obiwandsilva/passwordvalidator/application/web/http/controllers"
	"github.com/obiwandsilva/passwordvalidator/application/web/http/routes"
)

type PasswordValidator struct {
	EnvConfig                   config.EnvironmentConfig
	PasswordValidatorController *controllers.PasswordValidatorController
}

func NewPasswordValidator(
	envConfig config.EnvironmentConfig,
	passwordValidatorController *controllers.PasswordValidatorController,
) *PasswordValidator {
	return &PasswordValidator{
		EnvConfig:                   envConfig,
		PasswordValidatorController: passwordValidatorController,
	}
}

func (pw *PasswordValidator) Start() {
	router := gin.Default()

	routes.Routes(router, pw.PasswordValidatorController)

	web.NewServer(
		pw.EnvConfig.ServerPort,
		time.Second*time.Duration(pw.EnvConfig.ServerReadTimeout),
		time.Second*time.Duration(pw.EnvConfig.ServerWriteTimeout),
		router,
	).Start()
}
