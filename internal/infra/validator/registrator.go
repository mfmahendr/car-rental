package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Register(v *validator.Validate) {

	nikRegex := regexp.MustCompile(`^\d{13}$`)
	phoneRegex := regexp.MustCompile(`^(62|0)8\d{8,10}$`)

	v.RegisterValidation("carname", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[a-zA-Z ]+$`)
		return re.MatchString(fl.Field().String())
	})

	v.RegisterValidation("nik", func(fl validator.FieldLevel) bool {
		return nikRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return phoneRegex.MatchString(fl.Field().String())
	})
}
