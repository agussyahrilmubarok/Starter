package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error) map[string]string {
	errorsMap := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errorsMap
	}

	for _, fe := range validationErrors {
		field := strings.ToLower(fe.Field())
		switch fe.Tag() {
		case "required":
			errorsMap[field] = fmt.Sprintf("%s is required", capitalize(field))
		case "email":
			errorsMap[field] = "Invalid email format"
		case "unique":
			errorsMap[field] = fmt.Sprintf("%s already exists", capitalize(field))
		case "min":
			errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", capitalize(field), fe.Param())
		case "max":
			errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", capitalize(field), fe.Param())
		case "numeric":
			errorsMap[field] = fmt.Sprintf("%s must be a number", capitalize(field))
		default:
			errorsMap[field] = "Invalid value"
		}
	}

	return errorsMap
}

func capitalize(value string) string {
	if value == "" {
		return value
	}
	return strings.ToUpper(value[:1]) + value[1:]
}
