package controllers

import (
	"github.com/obiwandsilva/passwordvalidator/domain/interfaces/services"
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

func (pvc *PasswordValidatorController) ValidatePassword(w http.ResponseWriter, r *http.Request) {

}
