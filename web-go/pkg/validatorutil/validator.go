package validatorutil

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidatorError parses standard validation errors from the go-playground/validator library
// and maps them into a user-friendly map[string]string format.
// The key represents the struct field name, and the value is the custom error message.
func ValidatorError(err error) map[string]string {
	errorsMap := make(map[string]string)

	// Handler error from validator.v10
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()

			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "Invalid email format"
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf(
					"%s must be at least %s characters",
					field,
					fieldError.Param(),
				)
			case "max":
				errorsMap[field] = fmt.Sprintf(
					"%s must be at most %s characters",
					field,
					fieldError.Param(),
				)
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorsMap[field] = "Invalid value"
			}
		}
	}

	return errorsMap
}
