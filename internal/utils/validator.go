package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NoSpace(f string) bool {
	return !strings.Contains(f, " ")
}

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			filed := strings.ToLower(v.StructField())
			res[filed] = translateError(v)
		}
	}

	return res
}

func translateError(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "min":
		return fmt.Sprintf("%s is too short", fe.Field())
	case "max":
		return fmt.Sprintf("%s is too long", fe.Field())
	default:
		return fmt.Sprintf("%s is invalid", fe.Field())
	}

}

