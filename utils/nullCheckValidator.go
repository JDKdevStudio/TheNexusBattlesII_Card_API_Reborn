package utils

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func NullCheckValidator(obj interface{}, validate *validator.Validate) error {
	// Obtén el valor reflect.Value de la estructura.
	val := reflect.ValueOf(obj)

	// Verifica si es un puntero a una estructura y si es válido.
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("obj no es un puntero a una estructura válida")
	}

	// Obtén el valor reflect.Value de la estructura apuntada.
	val = val.Elem()

	// Obtén el tipo de la estructura.
	structType := val.Type()

	// Itera a través de los campos de la estructura.
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := structType.Field(i)

		// Si el campo no es nulo, realiza la validación.
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			fieldTag := fieldType.Tag.Get("validate")
			if fieldTag != "" {
				if err := validate.Var(field.Interface(), fieldTag); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
