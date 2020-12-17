package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/obiwandsilva/passwordvalidator/application/web/http/controllers"
)

func Routes(
	g *gin.Engine,
	passwordValidatorController *controllers.PasswordValidatorController,
) {
	g.POST("/validate", passwordValidatorController.ValidatePassword)
}