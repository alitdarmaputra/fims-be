package request

type HTTPUserCreateRequest struct {
	Email           string `json:"email"            binding:"required,email"`
	Password        string `json:"password"         binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Name            string `json:"name"             binding:"required"`
}

type HTTPUserLoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type HTTPChangePasswordRequest struct {
	NewPassword     string `json:"new_password"     binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

type HTTPResetTokenRequest struct {
	Email string `json:"email" binding:"required"`
}

type HTTPRedeemTokenRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
	Token       string `json:"token"        binding:"required"`
}

type VerificationParam struct {
	VerificationCode string `uri:"verification_code" binding:"required"`
}
