package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-challenge-exercise/internal/adapter/handler/http/dto"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type LoginHandler struct {
	service port.LoginService
}

func NewLoginHandler(service port.LoginService) *LoginHandler {
	return &LoginHandler{
		service: service,
	}
}

func (h *LoginHandler) UserLogin(ctx *gin.Context) {
	var request dto.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid body in login request binding"))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in login request validation", err))
		return
	}

	token, err := h.service.UserLogin(ctx, request.Email, request.Password)
	if err != nil {
		if pkg.IsBusinessError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in login request", err))
		return
	}

	ctx.JSON(http.StatusCreated, dto.LoginResponse{
		Token: token,
	})
}
