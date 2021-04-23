package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type M map[string]interface{}

func ResponseFormatter(code int, status string, message interface{}, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func ErrorFormatter(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
