package http

import (
	"net/http"
	"twitter-challenge-exercise/internal/adapter/handler/http/dto"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"

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
	// TODO LP: validaciones de datos

	testTweet := domain.Tweet{Text: "test"}

	testTweet, err := h.service.CreateTweet(ctx, testTweet)
	if err != nil {

	}

	ctx.JSON(http.StatusCreated, dto.MapTweetToTweetResponse(testTweet))
}
