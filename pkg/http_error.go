package pkg

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpError interface {
	GetStatusCode() int
	error
}

func ReturnHttpError(ctx *gin.Context, err HttpError) {
	// TODO LP: log!
	println(err.Error())

	ctx.JSON(err.GetStatusCode(), err)
}

// --------------------------------------------------------------------------------
type BadRequestError struct {
	Message string `json:"message"`
}

func NewBadRequestError(friendlyMessage string) HttpError {
	return BadRequestError{
		Message: friendlyMessage,
	}
}

func (e BadRequestError) GetStatusCode() int {
	return http.StatusBadRequest
}

func (e BadRequestError) Error() string {
	return e.Message
}

// --------------------------------------------------------------------------------
type RequestValidationError struct {
	Errors []string `json:"errors"`
}

func NewRequestValidationError(errors []string) HttpError {
	return RequestValidationError{
		Errors: errors,
	}
}

func (e RequestValidationError) GetStatusCode() int {
	return http.StatusBadRequest
}

func (e RequestValidationError) Error() string {
	return strings.Join(e.Errors, ", ")
}

// --------------------------------------------------------------------------------
type InternalServerError struct {
	Message string `json:"message"`
	error
}

func NewInternalServerError(friendlyMessage string, err error) HttpError {
	// Normally, we should log the internal error cause.
	// End user doesn't have to be aware of the internal error causes.
	return InternalServerError{
		Message: friendlyMessage,
		error:   err,
	}
}

func (e InternalServerError) GetStatusCode() int {
	return http.StatusInternalServerError
}

// --------------------------------------------------------------------------------
type ForbiddenError struct {
	Message string `json:"message"`
}

func NewForbiddenError(friendlyMessage string) HttpError {
	return ForbiddenError{
		Message: friendlyMessage,
	}
}

func (e ForbiddenError) GetStatusCode() int {
	return http.StatusForbidden
}

func (e ForbiddenError) Error() string {
	return e.Message
}
