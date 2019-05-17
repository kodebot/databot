package models

import (
	"time"

	"github.com/mmcdole/gofeed"
)

// Article model
type Article struct {
	Title         string
	PublishedDate time.Time
	Categories    []string
	ThumbImageURL string
	SourceURL     string
	Source        string
	OriginalFeed  gofeed.Item
	CreatedAt     time.Time
}
