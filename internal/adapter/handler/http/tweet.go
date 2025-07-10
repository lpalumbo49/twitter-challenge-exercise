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

type TweetHandler struct {
	service port.TweetService
}

func NewTweetHandler(service port.TweetService) *TweetHandler {
	return &TweetHandler{
		service: service,
	}
}

func (h *TweetHandler) CreateTweet(ctx *gin.Context) {
	var request dto.CreateTweetRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid body in create tweet request binding"))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in create tweet request validation", err))
		return
	}

	if request.UserID != ctx.GetUint64("user_id") {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("user_id not authorized"))
		return
	}

	tweet, err := h.service.CreateTweet(ctx, dto.MapCreateTweetRequestToTweet(request))
	if err != nil {
		if pkg.IsBusinessError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error creating new tweet", err))
		return
	}

	ctx.JSON(http.StatusCreated, dto.MapTweetToTweetResponse(tweet))
}

func (h *TweetHandler) UpdateTweet(ctx *gin.Context) {
	var request dto.UpdateTweetRequest

	idParam := ctx.Param("id")

	tweetID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid tweet id"))
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid body in update tweet request binding"))
		return
	}

	if err := pkg.ValidateStruct(request); err != nil {
		if ok, valErr := pkg.ParseStructValidationError(err); ok {
			pkg.ReturnHttpError(ctx, pkg.NewRequestValidationError(valErr.GetErrors()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error in update tweet request validation", err))
		return
	}

	if tweetID != request.ID {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("mismatch between tweet id and request tweet id"))
		return
	}

	if request.UserID != ctx.GetUint64("user_id") {
		pkg.ReturnHttpError(ctx, pkg.NewForbiddenError("user_id not authorized"))
		return
	}

	tweet, err := h.service.UpdateTweet(ctx, dto.MapUpdateTweetRequestToTweet(request))
	if err != nil {
		if !pkg.IsServerError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewBadRequestError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError("error updating tweet", err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapTweetToTweetResponse(tweet))
}

func (h *TweetHandler) GetTweetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	tweetID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewBadRequestError("invalid tweet id"))
		return
	}

	tweet, err := h.service.GetTweetByID(ctx, tweetID)
	if err != nil {
		if pkg.IsEntityNotFoundError(err) {
			pkg.ReturnHttpError(ctx, pkg.NewNotFoundError(err.Error()))
			return
		}

		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError(fmt.Sprintf("error searching for tweet id %d", tweetID), err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapTweetToTweetResponse(tweet))
}
