package dto

import "twitter-challenge-exercise/internal/core/domain"

type TimelineResponse struct {
	TweetResponse
	User UserResponse `json:"user"`
}

type TimelineResponses struct {
	Timeline []TimelineResponse `json:"timeline"`
}

func MapTimelineTweetToTimelineResponse(timelineTweet domain.TimelineTweet) TimelineResponse {
	return TimelineResponse{
		TweetResponse: MapTweetToTweetResponse(timelineTweet.Tweet),
		User:          MapUserToUserResponse(timelineTweet.User),
	}
}

func MapTimelineTweetsToTimelineResponses(timelineTweets []domain.TimelineTweet) TimelineResponses {
	responses := make([]TimelineResponse, 0)

	for _, timelineTweet := range timelineTweets {
		responses = append(responses, MapTimelineTweetToTimelineResponse(timelineTweet))
	}

	return TimelineResponses{
		Timeline: responses,
	}
}
