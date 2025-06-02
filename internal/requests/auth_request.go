package requests

type BasicAuth struct {
	Email    string `json:"email" validate:"required,email" example:"email@domain.com"`
	Password string `json:"password" validate:"required,min=5" example:"password"`
}

type RegisterRequest struct {
	BasicAuth
	Name  string `json:"name" validate:"required" example:"name"`
	Phone string `json:"phone" validate:"required" example:"081234567890"`
}

type LoginRequest struct {
	BasicAuth
}

type UpdateAccessTokenRequest struct {
	TokenId      uint   `json:"token_id" validate:"required" example:"1"`
	RefreshToken string `json:"refresh_token" validate:"required" example:"refresh_token"`
}
