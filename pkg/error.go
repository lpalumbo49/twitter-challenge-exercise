package pkg

import (
	"errors"
	"strings"
)

// --------------------------------------------------------------------------------
type StructValidationError struct {
	errors []string
}

func (e StructValidationError) Error() string {
	return strings.Join(e.errors, ", ")
}

func (e StructValidationError) GetErrors() []string {
	return e.errors
}

func IsStructValidationError(err error) bool {
	var valErr StructValidationError

	return errors.As(err, &valErr)
}

func ParseStructValidationError(err error) (bool, StructValidationError) {
	var valErr StructValidationError

	ok := errors.As(err, &valErr)
	return ok, valErr
}

func NewStructValidationError(errors []string) error {
	return StructValidationError{
		errors: errors,
	}
}

// --------------------------------------------------------------------------------
type EntityNotFoundError struct {
	message string
}

func (e EntityNotFoundError) Error() string {
	return e.message
}

func NewEntityNotFoundError(message string) error {
	return EntityNotFoundError{
		message: message,
	}
}

func IsEntityNotFoundError(err error) bool {
	var nfErr EntityNotFoundError

	return errors.As(err, &nfErr)
}

func ParseEntityNotFoundError(err error) (bool, EntityNotFoundError) {
	var nfErr EntityNotFoundError

	ok := errors.As(err, &nfErr)
	return ok, nfErr
}

// --------------------------------------------------------------------------------
type BusinessError struct {
	message string
}

func (e BusinessError) Error() string {
	return e.message
}

func NewBusinessError(message string) error {
	return BusinessError{
		message: message,
	}
}

func IsBusinessError(err error) bool {
	var nfErr BusinessError

	return errors.As(err, &nfErr)
}

func ParseBusinessError(err error) (bool, BusinessError) {
	var bErr BusinessError

	ok := errors.As(err, &bErr)
	return ok, bErr
}

// --------------------------------------------------------------------------------
type ServerError struct {
	message string
}

func (e ServerError) Error() string {
	return e.message
}

func NewServerError(message string) error {
	return ServerError{
		message: message,
	}
}

func IsServerError(err error) bool {
	var sErr ServerError

	return errors.As(err, &sErr)
}

func ParseServerError(err error) (bool, ServerError) {
	var sErr ServerError

	ok := errors.As(err, &sErr)
	return ok, sErr
}
