package utils

import (
	"github.com/go-playground/validator/v10"
)

func BoolValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		values := []bool{true, false}
		for _, item := range values {
			if item == fl.Field().Bool() {
				return true
			}
		}
		return false
	}
}
