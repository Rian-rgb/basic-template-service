package response

import (
	"basic-template-service/lib/utils"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MapValidationErrors(err error) []ValidationErrorField {
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return nil
	}

	var listErrors []ValidationErrorField

	for _, f := range validationErrors {
		var msg string

		fieldName := f.Field()

		switch f.Tag() {
		case "required":
			msg = fmt.Sprintf(
				"The %s field is required",
				utils.ToSnakeCase(fieldName),
			)

		case "email":
			msg = fmt.Sprintf(
				"The %s field must be a valid email address",
				utils.ToSnakeCase(fieldName),
			)

		case "min":
			msg = fmt.Sprintf(
				"The %s field must be at least %s characters long",
				utils.ToSnakeCase(fieldName),
				f.Param(),
			)

		case "max":
			msg = fmt.Sprintf(
				"The %s field must not exceed %s characters",
				utils.ToSnakeCase(fieldName),
				f.Param(),
			)

		case "numeric":
			msg = fmt.Sprintf(
				"The %s field must be numeric",
				utils.ToSnakeCase(fieldName),
			)

		default:
			msg = fmt.Sprintf(
				"The %s field failed validation for rule '%s'",
				utils.ToSnakeCase(fieldName),
				f.Tag(),
			)
		}

		listErrors = append(listErrors, ValidationErrorField{
			Field:   utils.ToSnakeCase(fieldName),
			Message: msg,
		})
	}

	return listErrors
}
