package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/data"
	"github.com/kodebot/newsfeed/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PruneArticles cleans old new items
func PruneArticles() {
	glog.Infoln("begining pruning articles...")
	defer handlePrunePanics()
	glog.Infoln("loading articles collection...")
	articlesCollection, err := data.GetCollection("articles")

	if err != nil {
		glog.Errorf("loading articles collections failed with error %s. pruning stopped.\n", err.Error())
		dumpErrorInDatabase("loading articles collections failed", errorDump{
			Error: err.Error()})
		return
	}

	glog.Infoln("pruning all the articles created 48 hours ago...")
	result, err := data.Delete(articlesCollection, bson.M{"createdat": bson.M{"$lte": time.Now().Add(-48 * time.Hour)}})

	if err != nil {
		glog.Errorf("deleting articles failed with error %s. pruning stopped.\n", err.Error())
		dumpErrorInDatabase("deleting articles failed", errorDump{
			Error: err.Error()})
		return
	}

	glog.Infof("deleted %d items", result.DeletedCount)

	glog.Infoln("pruning articles finished...")

}

// CreateArticles loads news articles into the database
func CreateArticles(data []map[string]interface{}) []*models.Article {
	glog.Info("creating articles...")
	result := []*models.Article{}
	totalItems := len(data)

	for i, item := range data {
		glog.Infof("creating article %d of %d\n", i+1, totalItems)

		article := models.Article{}

		for key, val := range item {
			if val != nil {
				fmt.Printf("%s, %+v\n", key, val)
			} else {
				fmt.Printf("%s, NIL\n", key)
			}
		}

		if title, ok := item["Title"].(string); ok {
			article.Title = title
		}

		if shortContent, ok := item["Description"].(string); ok {
			article.ShortContent = shortContent
		}

		if publishedDate, ok := item["PublishedDate"].(*time.Time); ok {
			article.PublishedDate = *publishedDate
		}

		if thumbImageURL, ok := item["ThumbImageUrl"].(string); ok {
			article.ThumbImageURL = thumbImageURL
		}

		if sourceURL, ok := item["SourceUrl"].(string); ok {
			article.SourceURL = sourceURL
		}

		article.CreatedAt = time.Now()

		result = append(result, &article)
		glog.Infof("finished creating article %d of %d\n", i+1, totalItems)
	}

	return result
}

// LoadArticles loads news articles into the database
func LoadArticles(articles []*models.Article) {
	glog.Infoln("loading articles collection...")
	articlesCollection, err := data.GetCollection("articles")
	if err != nil {
		if err != nil {
			glog.Errorf("loading articles collections failed with error %s. Skipping this source.\n", err.Error())
			dumpErrorInDatabase("loading articles collections failed", errorDump{
				Error: err.Error()})
			return
		}
	}
	glog.Infoln("loading articles collection finished...")

	totalItems := len(articles)

	for i, article := range articles {
		glog.Infof("loading item %d of %d\n", i+1, totalItems)
		glog.Infoln("checking whether the item has already been loaded...")
		var existing models.Article
		err = data.FindOne(articlesCollection, bson.M{"sourceurl": article.SourceURL}, &existing)

		if err != nil && err != mongo.ErrNoDocuments {
			glog.Errorf("error when checking if the item has already been loaded. error: %s\n", err.Error())
			dumpErrorInDatabase("error when checking if the item has already been loaded", errorDump{
				OtherData: article,
				Error:     err.Error()})
			continue
		}

		if existing.SourceURL != "" {
			glog.Infoln("item already found, not adding this to the store...")
			continue
		}

		glog.Infoln("item not found, adding this to the store...")
		var result *mongo.InsertOneResult
		result, err = data.InsertOne(articlesCollection, article)
		if err != nil {
			glog.Errorf("adding item to the store failed with error: %s. Skipping this item \n", err.Error())
			dumpErrorInDatabase("adding item to the store failed", errorDump{
				OtherData: article,
				Error:     err.Error()})
			continue
		}

		glog.Infof("item added to the store successfully. new id: %s\n", result)
		glog.Infof("finished loading item %d of %d\n", i+1, totalItems)
	}
}

func handleLoadFeedPanics(feedConfig models.FeedConfigItem) {
	defer panicInPanicHandler()
	if r := recover(); r != nil {
		glog.Infof("recovering from panic %s when loading feed from URL %s\n", r, feedConfig.URL)
		dumpErrorInDatabase("panic when loading feed", errorDump{FeedConfig: feedConfig, Error: r})
	}
}

func handlePrunePanics() {
	defer panicInPanicHandler()
	if r := recover(); r != nil {
		glog.Infof("recovering from panic %s when pruning\n", r)
		dumpErrorInDatabase("panic when pruning articles", errorDump{Error: r})
	}
}

func panicInPanicHandler() {
	if re := recover(); re != nil {
		glog.Infof("panic in panic handler %s the real error may not have been logged", re)
	}
}

func dumpErrorInDatabase(message string, errorDump errorDump) {
	defer func() {
		if r := recover(); r != nil {
			glog.Warningf("failed dumping error in the database. %s\n", r)
		}
	}()

	errors, err := data.GetCollection("errors")
	if err != nil {
		glog.Warningf("unable to load error collection to dump error. %s\n", err.Error())
		return
	}
	_, err = data.InsertOne(errors, errorDocument{message, errorDump, time.Now()})

	if err != nil {
		glog.Warningf("failed to insert error dump. %s\n", err.Error())
		return
	}
}

func convertTamilToEnglishDate(dateString string) (string, error) {
	tamilMonths := [...]string{"ஜனவரி", "பெப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"}
	englishMonths := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	for i, tamilMonth := range tamilMonths {
		dateString = strings.Replace(dateString, tamilMonth, englishMonths[i], -1)
	}
	dateString = strings.Replace(dateString, "  ", " ", -1) // remove double space
	return dateString, nil
}

type errorDocument struct {
	Message string
	Data    interface{}
	When    time.Time
}
type errorDump struct {
	FeedConfig models.FeedConfigItem
	OtherData  interface{}
	Error      interface{}
}
