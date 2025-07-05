package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	if err := validate.Struct(data); err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Field()+" is "+err.Tag())
		}
		return &ValidationError{Messages: errMsgs}
	}
	return nil
}

type ValidationError struct {
	Messages []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Messages, ", ")
}
