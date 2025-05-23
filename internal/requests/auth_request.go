package requests

type BasicAuth struct {
	Email    string `json:"email" validate:"required,email" example:"email@domain.com"`
	Password string `json:"password" validate:"required,min=5" example:"password"`
}

type Register struct {
	BasicAuth
	Name string `json:"name" validate:"required" example:"name"`
}

type LoginRequest struct {
	BasicAuth
}
