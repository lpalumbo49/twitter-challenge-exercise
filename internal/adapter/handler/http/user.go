package http

import (
	"fmt"
	"net/http"
	"strconv"
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
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid body in create user request binding"))
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
		if pkg.IsBusinessError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error creating new user", err))
		return
	}

	ctx.JSON(http.StatusCreated, dto.MapUserToUserResponse(user))
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var request dto.UpdateUserRequest

	idParam := ctx.Param("id")

	userID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid user_id"))
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid body in update user request binding"))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in update user request validation", err))
		return
	}

	if userID != request.ID {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("mismatch between user_id and request user_id"))
		return
	}

	if userID != ctx.GetUint64("user_id") {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("user_id not authorized"))
		return
	}

	user, err := h.service.UpdateUser(ctx, dto.MapUpdateUserRequestToUser(request))
	if err != nil {
		if !pkg.IsServerError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error updating user", err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapUserToUserResponse(user))
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	userID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid user id"))
		return
	}

	user, err := h.service.GetUserByID(ctx, userID)
	if err != nil {
		if pkg.IsEntityNotFoundError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewNotFoundError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError(fmt.Sprintf("error searching for user id %d", userID), err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapUserToUserResponse(user))
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.SearchUsers(ctx)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error searching for users", err))
		return
	}

	ctx.JSON(http.StatusOK, dto.UsersResponse{
		Users: dto.MapUsersToUserResponses(users),
	})
}
