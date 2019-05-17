package services

import (
	"reflect"
	"regexp"
	"time"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/data"
	"github.com/kodebot/newsfeed/models"
	"github.com/mmcdole/gofeed"
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

// LoadArticlesFromFeed from the given url
func LoadArticlesFromFeed(feedConfig models.FeedConfigItem) {
	defer handleLoadFeedPanics(feedConfig)
	glog.Infof("starting loading articles from URL: %s \n", feedConfig.URL)
	defer glog.Infof("ending loading feed from URL: %s", feedConfig.URL)

	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedConfig.URL)
	if err != nil {
		glog.Errorf("parsing feed failed with error %s. Skipping this source.\n", err.Error())
		dumpErrorInDatabase("parsing feed failed", errorDump{
			FeedConfig: feedConfig,
			Error:      err.Error()})
		return
	}

	totalItems := len(feed.Items)
	if totalItems == 0 {
		glog.Infoln("no items found to process")
		return
	}

	glog.Infof("%d items found\n", totalItems)

	glog.Infoln("loading articles collection...")
	articlesCollection, err := data.GetCollection("articles")
	if err != nil {
		if err != nil {
			glog.Errorf("loading articles collections failed with error %s. Skipping this source.\n", err.Error())
			dumpErrorInDatabase("loading articles collections failed", errorDump{
				FeedConfig: feedConfig,
				Error:      err.Error()})
			return
		}
	}
	glog.Infoln("loading articles collection...")

	for i, item := range feed.Items {
		glog.Infof("processing item %d of %d\n", i+1, totalItems)

		glog.Infof("item data dump %+v\n", item)

		var revisedDate time.Time
		if item.PublishedParsed != nil {
			revisedDate = *item.PublishedParsed
			if isUtcMidnight(item.PublishedParsed) { // this means we don't have time - only date is present - just use current time
				glog.Infoln("published date missing time, setting it current time...")
				revisedDate = time.Now()
			}
		} else {
			glog.Infoln("parsed date not found. use current date time...")
			revisedDate = time.Now()
		}

		glog.Infoln("extracting source url...")
		extractedSourceURL, err := extractData(item, feedConfig.ItemURLExtractor)
		glog.Infof("extracted source url %s \n", extractedSourceURL)

		if err != nil {
			glog.Errorf("extracting source link failed with error %s. Skipping this item.\n", err.Error())
			dumpErrorInDatabase("extracting source link failed", errorDump{
				FeedConfig: feedConfig,
				OtherData:  item,
				Error:      err.Error()})
			continue
		}

		if extractedSourceURL == "" {
			glog.Errorln("extracting source link failed - no source link found. Skipping this item.")
			dumpErrorInDatabase("extracting source link failed", errorDump{
				FeedConfig: feedConfig,
				OtherData:  item,
				Error:      err.Error()})
			continue
		}

		glog.Infoln("extracting thumb image url...")
		extactedImageURL, err := extractData(item, feedConfig.ItemThumbImageExtractor)
		glog.Infof("extracted thumb image url %s \n", extactedImageURL)

		if err != nil {
			glog.Warningf("extracting image link failed with error: %s\n", err.Error())
		}

		if extactedImageURL == "" {
			glog.Warningln("extracting image link failed - no image link found.")
		}

		article := models.Article{
			Title:         item.Title,
			PublishedDate: revisedDate,
			Categories:    []string{feedConfig.DefaultCategory},
			ThumbImageURL: extactedImageURL,
			SourceURL:     extractedSourceURL,
			Source:        feedConfig.Origin,
			OriginalFeed:  *item,
			CreatedAt:     time.Now()}

		glog.Infoln("preparing to store the item...")

		glog.Infoln("checking whether the item has already been loaded...")
		var existing models.Article
		err = data.FindOne(articlesCollection, bson.M{"sourceurl": extractedSourceURL}, &existing)

		if err != nil && err != mongo.ErrNoDocuments {
			glog.Errorf("error when checking if the item has already been loaded. error: %s\n", err.Error())
			dumpErrorInDatabase("error when checking if the item has already been loaded", errorDump{
				FeedConfig: feedConfig,
				OtherData:  item,
				Error:      err.Error()})
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
				FeedConfig: feedConfig,
				OtherData:  item,
				Error:      err.Error()})
			continue
		}

		glog.Infof("item added to the store successfully. new id: %s\n", result)
		glog.Infof("finished processing item %d of %d\n", i+1, totalItems)
	}

}

func extractData(source interface{}, extractorConfig models.FeedDataExtractorConfig) (string, error) {
	glog.Infoln("extracting data using extractor...")
	result := ""
	rawSourceData := reflect.Indirect(reflect.ValueOf(source)).FieldByName(extractorConfig.SourceField).String()
	glog.Infof("raw data dump: %s\n", rawSourceData)
	if extractorConfig.ScrapingRequired != true {
		glog.Infoln("scraping is not required, returning the raw data...")
		return rawSourceData, nil
	}

	glog.Infoln("scraping required...")
	switch extractorConfig.SelectorType {
	case "regexp":
		glog.Infoln("scarping using regexp selector type...")
		re, err := regexp.Compile(extractorConfig.Selector)
		if err != nil {
			glog.Errorf("invalid regexp: %s error: %s. \n", extractorConfig.Selector, err.Error())
			return "", err
		}

		requiredMatchIndex := 0
		for i, val := range re.SubexpNames() {
			if val == "URL" {
				requiredMatchIndex = i
			}
		}

		if requiredMatchIndex == 0 {
			glog.Errorf("invalid regexp: %s no named group called URL is found. \n", extractorConfig.Selector)
			return "", err
		}
		matches := re.FindStringSubmatch(rawSourceData)
		if len(matches) < requiredMatchIndex+1 {
			glog.Warningln("no match found.")
			return "", err
		}

		return matches[requiredMatchIndex], nil
		// todo: other cases like xpath and css are not supported yet
	}
	return result, nil
}

func isUtcMidnight(datetime *time.Time) bool {
	glog.Infoln("checking if the date is midnight utc...")
	return (datetime.Hour() == 0 &&
		datetime.Minute() == 0 &&
		datetime.Second() == 0 &&
		datetime.Nanosecond() == 0 &&
		datetime.Location().String() == "UTC")
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
