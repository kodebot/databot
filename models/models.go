package models

import (
	"time"

	"github.com/mmcdole/gofeed"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ArticleMinimal used as DTO for api
type ArticleMinimal struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string
	PublishedDate time.Time
	ThumbImageURL string
	Source        string
}

// ArticleContent DTO to show the detail of the new article
type ArticleContent struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string
	PublishedDate time.Time
	ThumbImageURL string
	ShortContent  string
	SourceURL     string
	Source        string
}

// Article model
type Article struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string
	ShortContent  string
	PublishedDate time.Time
	Categories    []string
	ThumbImageURL string
	SourceURL     string
	Source        string
	OriginalFeed  gofeed.Item
	CreatedAt     time.Time
}
