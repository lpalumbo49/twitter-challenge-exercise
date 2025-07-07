package pkg

import (
	"errors"
	"strings"
)

// TODO LP: switch con las conversiones de errores

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
type NotFoundError struct {
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(message string) error {
	return NotFoundError{
		message: message,
	}
}

func IsNotFoundError(err error) bool {
	var nfErr NotFoundError

	return errors.As(err, &nfErr)
}

func ParseNotFoundError(err error) (bool, NotFoundError) {
	var nfErr NotFoundError

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
