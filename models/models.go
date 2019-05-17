package models

import (
	"time"

	"github.com/mmcdole/gofeed"
)

// ArticleMinimal used as DTO for api
type ArticleMinimal struct {
	Title         string
	PublishedDate time.Time
	Categories    []string
	ThumbImageURL string
	SourceURL     string
	Source        string
}

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
