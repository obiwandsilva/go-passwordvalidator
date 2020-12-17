package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/obiwandsilva/passwordvalidator/application/web/http/requests"
	"github.com/obiwandsilva/passwordvalidator/domain/interfaces/services"
	"log"
	"net/http"
)

type PasswordValidatorController struct {
	passwordValidatorService services.PasswordValidatorService
}

func NewPasswordValidatorController(
	passwordValidatorService services.PasswordValidatorService,
) *PasswordValidatorController {
	return &PasswordValidatorController{
		passwordValidatorService: passwordValidatorService,
	}
}

func (pvc *PasswordValidatorController) ValidatePassword(c *gin.Context) {
	var validatePasswordRequest requests.ValidatePasswordRequest

	err := c.ShouldBindJSON(&validatePasswordRequest)
	if err != nil {
		log.Printf("error when decoding request body: %v\n", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	validation := pvc.passwordValidatorService.IsValid(validatePasswordRequest.Password)
	c.JSON(http.StatusOK, validation)
}
