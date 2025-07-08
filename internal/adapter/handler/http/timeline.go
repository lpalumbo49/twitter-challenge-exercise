package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-challenge-exercise/internal/adapter/handler/http/dto"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
)

type TimelineHandler struct {
	service port.TimelineService
}

func NewTimelineHandler(service port.TimelineService) *TimelineHandler {
	return &TimelineHandler{
		service: service,
	}
}

func (h *TimelineHandler) GetTimelineByUserID(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	timeline, err := h.service.GetTimelineByUserID(ctx, userID)
	if err != nil {
		pkg.ReturnHttpError(ctx, pkg.NewInternalServerError(fmt.Sprintf("error searching for timeline for user %d", userID), err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapTimelineTweetsToTimelineResponses(timeline))
}
