package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-challenge-exercise/internal/adapter/handler/http/dto"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type FollowerHandler struct {
	service port.FollowerService
}

func NewFollowerHandler(service port.FollowerService) *FollowerHandler {
	return &FollowerHandler{
		service: service,
	}
}

func (h *FollowerHandler) CreateFollower(ctx *gin.Context) {
	var request dto.CreateFollowerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in create follower request binding", err))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in create follower request validation", err))
		return
	}

	if request.UserID != ctx.GetUint64("user_id") {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("user_id not authorized"))
		return
	}

	follower, err := h.service.CreateFollower(ctx, dto.MapCreateFollowerRequestToFollower(request))
	if err != nil {
		if pkg.IsBusinessError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error creating new follower", err))
		return
	}

	ctx.JSON(http.StatusCreated, dto.MapFollowerToFollowerResponse(follower))
}
