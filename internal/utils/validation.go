package utils

import (
	"github.com/go-playground/validator/v10"
	r "github.com/hnifmaghfur/go-user-service/internal/requests"
)

var validate = validator.New()

func ValidateLoginRequest(loginRequest r.LoginRequest) error {
	if err := validate.Struct(loginRequest); err != nil {
		return err
	}
	return nil
}
