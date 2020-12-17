package services

type PasswordValidatorService interface {
	IsValid(password string) (bool, []string)
}
