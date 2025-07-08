package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg"
	"twitter-challenge-exercise/pkg/mysql"
)

const (
	insertTweetQuery     = "INSERT INTO tweet(user_id, text, created_at, updated_at) VALUES(?, ?, ?, ?)"
	updateTweetQuery     = "UPDATE tweet SET text = ?, updated_at = ? WHERE id = ?"
	selectTweetByIDQuery = "SELECT id, user_id, text, created_at, updated_at FROM tweet WHERE id = ?"
)

type tweetRepository struct {
	db *mysql.DB
}

func NewTweetRepository(db *mysql.DB) port.TweetRepository {
	return &tweetRepository{
		db: db,
	}
}

func (t *tweetRepository) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	now := time.Now()

	result, err := t.db.Exec(insertTweetQuery, tweet.UserID, tweet.Text, now, now)
	if err != nil {
		return tweet, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return tweet, err
	}

	tweet.ID = uint64(lastInsertedID)
	tweet.CreatedAt = now
	tweet.UpdatedAt = now

	return tweet, nil
}

func (t *tweetRepository) UpdateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	now := time.Now()

	_, err := t.db.Exec(updateTweetQuery, tweet.Text, now, tweet.ID)
	if err != nil {
		return tweet, err
	}

	tweet.UpdatedAt = now

	return tweet, nil
}

func (t *tweetRepository) GetTweetByID(ctx context.Context, tweetID uint64) (domain.Tweet, error) {
	var tweet domain.Tweet

	row := t.db.QueryRow(selectTweetByIDQuery, tweetID)

	err := row.Scan(&tweet.ID, &tweet.UserID, &tweet.Text, &tweet.CreatedAt, &tweet.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tweet, pkg.NewEntityNotFoundError(fmt.Sprintf("tweet_id %d not found", tweetID))
		}

		return tweet, err
	}

	return tweet, nil
}
