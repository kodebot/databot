package models

import (
	"time"

	"github.com/mmcdole/gofeed"
)

// NewsItems model
type NewsItem struct {
	Title         string
	PublishedDate time.Time
	Categories    []string
	ThumbImageURL string
	SourceURL     string
	Source        string
	OriginalFeed  gofeed.Item
	CreatedAt     time.Time
}
