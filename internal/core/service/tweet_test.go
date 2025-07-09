package service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/service"
	"twitter-challenge-exercise/pkg"
)

func TestTweet_CreateTweet_Success(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("CreateTweet", mock.Anything, testTweet).Return(testTweet, nil)
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil)

	tweetService := service.NewTweetService(tweetRepository)
	responseTweet, err := tweetService.CreateTweet(context.Background(), testTweet)

	assert.NoError(t, err)
	assert.Equal(t, testTweet, responseTweet)

	tweetRepository.AssertExpectations(t)
}

func TestTweet_CreateTweet_Error(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("CreateTweet", mock.Anything, testTweet).Return(testTweet, errors.New("error in database"))

	tweetService := service.NewTweetService(tweetRepository)
	_, err := tweetService.CreateTweet(context.Background(), testTweet)

	assert.Error(t, err)
	assert.True(t, pkg.IsServerError(err))

	tweetRepository.AssertExpectations(t)
}

func TestTweet_UpdateTweet_Success(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil).Once()

	testNewTweet := getTestTweet()
	testNewTweet.Text = "Oops, fixing some typo in tweet"

	tweetRepository.On("GetTweetByID", mock.Anything, testNewTweet.ID).Return(testNewTweet, nil).Once()
	tweetRepository.On("UpdateTweet", mock.Anything, testNewTweet).Return(testNewTweet, nil)

	tweetService := service.NewTweetService(tweetRepository)

	_, err := tweetService.UpdateTweet(context.Background(), testNewTweet)
	assert.NoError(t, err)

	tweetRepository.AssertExpectations(t)
}

func TestTweet_UpdateTweet_DifferentUserID(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil)

	tweetService := service.NewTweetService(tweetRepository)

	testTweet.UserID = 55
	_, err := tweetService.UpdateTweet(context.Background(), testTweet)

	assert.Error(t, err, "user 55 does not own this tweet")

	tweetRepository.AssertExpectations(t)
}

func TestTweet_UpdateTweet_ErrorGettingExistingTweet(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, errors.New("error in database"))

	tweetService := service.NewTweetService(tweetRepository)
	_, err := tweetService.UpdateTweet(context.Background(), testTweet)

	assert.Error(t, err)

	tweetRepository.AssertExpectations(t)
}

func TestTweet_UpdateTweet_Error(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil)

	testTweet.Text = "Oops, fixing some typo in tweet"
	tweetRepository.On("UpdateTweet", mock.Anything, testTweet).Return(testTweet, errors.New("error in database"))

	tweetService := service.NewTweetService(tweetRepository)

	_, err := tweetService.UpdateTweet(context.Background(), testTweet)

	assert.Error(t, err)
	assert.True(t, pkg.IsServerError(err))

	tweetRepository.AssertExpectations(t)
}

func TestTweet_GetTweetByID_Success(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, nil)

	tweetService := service.NewTweetService(tweetRepository)

	tweet, err := tweetService.GetTweetByID(context.Background(), testTweet.ID)

	assert.NoError(t, err)
	assert.Equal(t, testTweet, tweet)

	tweetRepository.AssertExpectations(t)
}

func TestTweet_GetTweetByID_Error(t *testing.T) {
	testTweet := getTestTweet()

	tweetRepository := database.NewTweetMockRepository()
	tweetRepository.On("GetTweetByID", mock.Anything, testTweet.ID).Return(testTweet, errors.New("error in database"))

	tweetService := service.NewTweetService(tweetRepository)

	_, err := tweetService.GetTweetByID(context.Background(), testTweet.ID)

	assert.Error(t, err, "error in database")

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
