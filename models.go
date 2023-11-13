package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}


type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: feed.LastFetchedAt,
	}
}

type FeedFollow struct {
    ID        uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    FeedID    uuid.UUID `json:"feed_id"`
    UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedFollowToFeedFollow(feedFollow FeedFollow) FeedFollow {
    return FeedFollow{
        ID:        feedFollow.ID,
        CreatedAt: feedFollow.CreatedAt,
        UpdatedAt: feedFollow.UpdatedAt,
        FeedID:    feedFollow.FeedID,
        UserID:    feedFollow.UserID,
    }
}
func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
