package database

import (
	"context"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/port"
	"twitter-challenge-exercise/pkg/mysql"
)

const (
	selectTimelineByUserIDQuery = `
		SELECT t.id, t.user_id, t.text, t.created_at, t.updated_at, u.id, u.name, u.surname, u.email, u.username,
		       u.created_at, u.updated_at
		FROM tweet t
			INNER JOIN follower f ON t.user_id = f.followed_by_user_id
			INNER JOIN user u ON u.id = t.user_id
		WHERE f.user_id = ?
		ORDER BY t.created_at DESC
`
)

type timelineRepository struct {
	db *mysql.DB
}

func NewTimelineRepository(db *mysql.DB) port.TimelineRepository {
	return &timelineRepository{
		db: db,
	}
}

func (t *timelineRepository) GetTimelineByUserID(ctx context.Context, userID uint64) ([]domain.TimelineTweet, error) {
	var timeline []domain.TimelineTweet

	rows, err := t.db.Query(selectTimelineByUserIDQuery, userID)
	if err != nil {
		return timeline, err
	}

	defer rows.Close()
	for rows.Next() {
		var timelineTweet domain.TimelineTweet

		err = rows.Scan(&timelineTweet.ID, &timelineTweet.UserID, &timelineTweet.Text, &timelineTweet.CreatedAt,
			&timelineTweet.UpdatedAt, &timelineTweet.User.ID, &timelineTweet.User.Name, &timelineTweet.User.Surname,
			&timelineTweet.User.Email, &timelineTweet.User.Username, &timelineTweet.User.CreatedAt, &timelineTweet.User.UpdatedAt)
		if err != nil {
			return timeline, err
		}

		timeline = append(timeline, timelineTweet)
	}

	return timeline, nil
}
