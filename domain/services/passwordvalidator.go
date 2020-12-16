package services

type PasswordValidator struct {}

func NewPasswordValidator() *PasswordValidator {
	return &PasswordValidator{}
}

func (pvs PasswordValidator) IsValid(password string) bool {
	return true
}
