package services

type PasswordValidator struct {}

func (pvs PasswordValidator) IsValid(password string) bool {
	return true
}
