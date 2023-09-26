package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegexValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		validateTag := fl.Param()
		value := fl.Field().String()
		regex := regexp.MustCompile(validateTag)
		return regex.MatchString(value)
	}
}
