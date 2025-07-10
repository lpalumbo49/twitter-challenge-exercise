package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	handler "twitter-challenge-exercise/internal/adapter/handler/http"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/service"
)

func TestTweetHandler_UpdateTweet_InvalidRequestBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := handler.NewTweetHandler(tweetService)

	router := gin.Default()
	router.PUT("/api/v1/tweet/:id", tweetHandler.UpdateTweet)

	requestBody := map[string]interface{}{
		"id":      "49", // Number, but string
		"user_id": testTweet.UserID,
		"text":    testTweet.Text,
	}
	jsonBody, _ := json.Marshal(requestBody)

	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/tweet/%d", testTweet.ID), bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var responseMap map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseMap)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "invalid body in update tweet request binding", responseMap["message"])

	tweetRepository.AssertExpectations(t)
}

func TestTweetHandler_UpdateTweet_TweetIDMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := handler.NewTweetHandler(tweetService)

	router := gin.Default()
	router.PUT("/api/v1/tweet/:id", tweetHandler.UpdateTweet)

	requestBody := map[string]interface{}{
		"id":      testTweet.ID,
		"user_id": testTweet.UserID,
		"text":    testTweet.Text,
	}
	jsonBody, _ := json.Marshal(requestBody)

	anotherTweetID := 55

	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/tweet/%d", anotherTweetID), bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var responseMap map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseMap)

	assert.Equal(t, http.StatusForbidden, response.Code)
	assert.Equal(t, "mismatch between tweet id and request tweet id", responseMap["message"])

	tweetRepository.AssertExpectations(t)
}

func TestTweetHandler_UpdateTweet_UserIDNotAuthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := handler.NewTweetHandler(tweetService)

	response := httptest.NewRecorder()
	_, router := gin.CreateTestContext(response)

	router.Use(func(c *gin.Context) {
		// Supposing that the user is already authenticated (and a different one from this tweet)
		c.Set("user_id", uint64(55))
	})

	router.PUT("/api/v1/tweet/:id", tweetHandler.UpdateTweet)

	requestBody := map[string]interface{}{
		"id":      testTweet.ID,
		"user_id": testTweet.UserID,
		"text":    testTweet.Text,
	}
	jsonBody, _ := json.Marshal(requestBody)

	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/tweet/%d", testTweet.ID), bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	var responseMap map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseMap)

	assert.Equal(t, http.StatusForbidden, response.Code)
	assert.Equal(t, "user_id not authorized", responseMap["message"])

	tweetRepository.AssertExpectations(t)
}

func TestTweetHandler_UpdateTweet_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil).Once()

	testNewTweet := getTestTweet()
	testNewTweet.Text = "Oops, fixing some typo in tweet"

	tweetRepository.On("GetTweetByID", mock.Anything, testNewTweet.ID).Return(testNewTweet, nil).Once()
	tweetRepository.On("UpdateTweet", mock.Anything, testNewTweet).Return(testNewTweet, nil)

	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := handler.NewTweetHandler(tweetService)

	response := httptest.NewRecorder()
	_, router := gin.CreateTestContext(response)

	router.Use(func(c *gin.Context) {
		// Supposing that the user is already authenticated
		c.Set("user_id", testTweet.UserID)
	})

	router.PUT("/api/v1/tweet/:id", tweetHandler.UpdateTweet)

	requestBody := map[string]interface{}{
		"id":      testNewTweet.ID,
		"user_id": testNewTweet.UserID,
		"text":    testNewTweet.Text,
	}
	jsonBody, _ := json.Marshal(requestBody)

	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/tweet/%d", testNewTweet.ID), bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	var responseMap map[string]interface{}
	_ = json.Unmarshal(response.Body.Bytes(), &responseMap)

	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, float64(testNewTweet.ID), responseMap["id"])
	assert.Equal(t, float64(testNewTweet.UserID), responseMap["user_id"])
	assert.Equal(t, testNewTweet.Text, responseMap["text"])
	assert.NotNil(t, testNewTweet.CreatedAt, responseMap["created_at"])
	assert.NotNil(t, testNewTweet.UpdatedAt, responseMap["updated_at"])

	tweetRepository.AssertExpectations(t)
}

func getTestTweet() domain.Tweet {
	return domain.Tweet{
		ID:        49,
		UserID:    42,
		Text:      "This is some kind of text",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
