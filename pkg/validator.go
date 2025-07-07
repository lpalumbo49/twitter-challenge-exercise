package pkg

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()

	// Return 'json' tag in field name, not the struct value name
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateStruct(s any) error {
	var validationErrors validator.ValidationErrors

	err := v.Struct(s)
	if err == nil {
		return nil
	}

	if errors.As(err, &validationErrors) {
		var errs []string

		for _, validErr := range validationErrors {
			switch validErr.Tag() {
			case "required":
				errs = append(errs, fmt.Sprintf("field '%s' is required", validErr.Field()))
			case "min":
				errs = append(errs, fmt.Sprintf("field '%s' should have a minimum length of %s", validErr.Field(), validErr.Param()))
			case "max":
				errs = append(errs, fmt.Sprintf("field '%s' has exceeded maximum length of %s", validErr.Field(), validErr.Param()))
			case "email":
				errs = append(errs, fmt.Sprintf("field '%s' has an invalid email format", validErr.Field()))
			default:
				errs = append(errs, fmt.Sprintf("field '%s' failed for the '%s' constraint", validErr.Field(), validErr.Tag()))
			}
		}

		return NewStructValidationError(errs)
	}

	return err
}
