package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func APIResponse(message string, code int, status string, data any) Response {
	return Response{
		Status:  status,
		Message: status,
		Data:    data,
	}
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
