package services

import "github.com/obiwandsilva/passwordvalidator/domain/entities"

type PasswordValidatorService interface {
	IsValid(password string) entities.Validation
}
