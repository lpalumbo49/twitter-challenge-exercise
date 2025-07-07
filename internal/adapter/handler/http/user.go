package http

import (
	"net/http"
	"twitter-challenge-exercise/internal/adapter/handler/http/dto"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in create user request binding", err))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in create user request validation", err))
		return
	}

	user, err := h.service.CreateUser(ctx, dto.MapCreateUserRequestToUser(request))
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error creating new user", err))
		return
	}

	ctx.JSON(http.StatusCreated, dto.MapUserToCreateUserResponse(user))
}

// TODO LP: el resto de los métodos
