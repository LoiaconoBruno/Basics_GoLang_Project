package main // O el paquete correspondiente

import (
	"time"

	"coursegolang/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"user-id"`
	Name    string    `json:"user-name"`
	Email   string    `json:"user-email"`
	Created time.Time `json:"user-created"`
	APIKey  string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:      dbUser.ID,
		Name:    dbUser.Name,
		Email:   dbUser.Email,
		Created: dbUser.Created,
		APIKey:  dbUser.ApiKey,
	}
}

type Feed struct {
	IDFeeds   uuid.UUID `json:"feed-id"`
	CreatedAt time.Time `json:"feed-created-at"`
	Name      string    `json:"feed-name"`
	Url       string    `json:"feed-url"`
	UserID    uuid.UUID `json:"feed-user-id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		IDFeeds:   dbFeed.IDFeeds,
		CreatedAt: dbFeed.CreatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeed {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}

	return feeds
}
