package middlewares

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func InitValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
