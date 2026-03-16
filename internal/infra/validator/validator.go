package validator

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	Register(v)
	return v
}
