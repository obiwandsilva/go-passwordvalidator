package requests

type ValidatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}
