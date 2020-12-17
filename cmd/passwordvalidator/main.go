package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/obiwandsilva/passwordvalidator/application"
	"github.com/obiwandsilva/passwordvalidator/application/config"
	"github.com/obiwandsilva/passwordvalidator/application/web/http/controllers"
	"github.com/obiwandsilva/passwordvalidator/domain/services"
	"log"
)

func main() {
	envConfig := config.EnvironmentConfig{}
	if err := env.Parse(&envConfig); err != nil {
		log.Panicf("error when loading environment configuration: %v", err)
	}

	passwordValidatorService := services.NewPasswordValidator(envConfig)
	passwordValidatorController := controllers.NewPasswordValidatorController(passwordValidatorService)
	passwordValidatorApplication := application.NewPasswordValidator(envConfig, passwordValidatorController)

	log.Printf("Running Server on port %s...", envConfig.ServerPort)
	passwordValidatorApplication.Start()
}
