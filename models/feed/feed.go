package feed

import (
	"time"

	"github.com/google/uuid"
	"github.com/rayhan889/rss-aggr/internal/database"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	UserID  uuid.UUID `json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func HandleFeedToFeedCustomModel(dbFeed database.Feed) Feed {
	return Feed {
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func HandleFeedsToFeedsCustomModel(dbFeeds []database.Feed) []Feed {
	sliceFeeds := make([]Feed, len(dbFeeds), cap(dbFeeds))
	for i, dbFeed := range dbFeeds {
		sliceFeeds[i] = HandleFeedToFeedCustomModel(dbFeed)
	}
	return sliceFeeds
}