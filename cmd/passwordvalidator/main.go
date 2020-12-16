package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/obiwandsilva/passwordvalidator/appplication"
	"github.com/obiwandsilva/passwordvalidator/appplication/config"
	"github.com/obiwandsilva/passwordvalidator/appplication/web/http/controllers"
	"github.com/obiwandsilva/passwordvalidator/domain/services"
	"log"
)

func main() {
	envConfig := config.EnvironmentConfig{}
	if err := env.Parse(&envConfig); err != nil {
		log.Panicf("error when loading environment configuration: %v", err)
	}

	passwordValidatorService := services.NewPasswordValidator()
	passwordValidatorController := controllers.NewPasswordValidatorController(passwordValidatorService)
	passwordValidatorApplication := appplication.PasswordValidator{
		EnvConfig:                   envConfig,
		PasswordValidatorController: passwordValidatorController,
	}

	passwordValidatorApplication.Start()
}
