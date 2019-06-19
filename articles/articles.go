package articles

import (
	"time"

	"github.com/kodebot/newsfeed/logger"

	"github.com/kodebot/newsfeed/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Article model
type Article struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string
	ShortContent  string
	Content       string
	PublishedDate time.Time
	Categories    []string
	ThumbImageURL string
	SourceURL     string
	Source        string
	OriginalFeed  interface{}
	CreatedAt     time.Time
}

// NewArticle creates a new Article
func NewArticle(item map[string]interface{}) *Article {
	article := Article{}

	if title, ok := item["Title"].(string); ok {
		article.Title = title
	}

	if shortContent, ok := item["Description"].(string); ok {
		article.ShortContent = shortContent
	}

	if publishedDate, ok := item["PublishedDate"].(*time.Time); ok {
		article.PublishedDate = *publishedDate
	} else if publishedDate, ok := item["PublishedDate"].(time.Time); ok {
		article.PublishedDate = publishedDate
	}

	if thumbImageURL, ok := item["ThumbImageUrl"].(string); ok {
		article.ThumbImageURL = thumbImageURL
	}

	if sourceURL, ok := item["SourceUrl"].(string); ok {
		article.SourceURL = sourceURL
	}

	if item["Category"] != nil {
		if category, ok := item["Category"].(string); ok {
			article.Categories = append(article.Categories, category)
		}
	}
	article.CreatedAt = time.Now()

	return &article
}

// Store adds/updates Article in the database
func (article *Article) Store(articleCollection *mongo.Collection) error {
	var existing Article
	err := data.FindOne(articleCollection, bson.M{"sourceurl": article.SourceURL}, &existing)

	if err != nil && err != mongo.ErrNoDocuments {
		logger.Errorf("error when checking if the item has already been loaded. error: %s\n", err.Error())
		return err
	}

	if existing.SourceURL != "" {
		// todo: update when the item is already found
		logger.Infof("item already found, not adding this to the store...")
		return nil
	}

	logger.Infof("item not found, adding this to the store...")
	var result *mongo.InsertOneResult
	result, err = data.InsertOne(articleCollection, article)
	if err != nil {
		logger.Errorf("adding item to the store failed with error: %s. Skipping this item \n", err.Error())
		return err
	}

	logger.Infof("item added to the store successfully. new id: %s\n", result)
	return nil
}
