package internalErrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func ValidStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validatorErrors := err.(validator.ValidationErrors)
	fieldError := validatorErrors[0]

	return errors.New(fieldError.StructField() + " " + fieldError.Tag() + " " + fieldError.Param())
}
