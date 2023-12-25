package validations

import "github.com/go-playground/validator/v10"

func GetValidator() *validator.Validate {
	validate := validator.New()

	return validate
}
