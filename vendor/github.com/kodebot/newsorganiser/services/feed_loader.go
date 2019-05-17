package services

import (
	"reflect"
	"regexp"
	"time"

	"github.com/golang/glog"
	"github.com/kodebot/newsorganiser/data"
	"github.com/kodebot/newsorganiser/models"
	"github.com/mmcdole/gofeed"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadFeed from the given url
func LoadFeed(feedConfig models.FeedConfigItem) {

	glog.Infof("starting loading feed from URL: %s \n", feedConfig.URL)
	defer glog.Infof("ending loading feed from URL: %s", feedConfig.URL)

	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedConfig.URL)
	if err != nil {
		glog.Errorf("parsing feed failed with error %s. Skipping this source.\n", err.Error())
		return
	}

	totalItems := len(feed.Items)
	if totalItems == 0 {
		glog.Infoln("no items found to process")
		return
	}

	glog.Infof("%d items found\n", totalItems)

	glog.Infoln("loading newsitems collection...")
	newsItemCollection, err := data.GetCollection("newsitems")
	if err != nil {
		if err != nil {
			glog.Errorf("loading newsitems collections failed with error %s. Skipping this source.\n", err.Error())
			return
		}
	}
	glog.Infoln("loading newsitems collection...")

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
			glog.Errorf("extracting source link failed with error %s\n. Skipping this item.", err.Error())
			continue
		}

		if extractedSourceURL == "" {
			glog.Errorln("extracting source link failed - no source link found. Skipping this item.")
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

		newsItem := models.NewsItem{
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
		var existing models.NewsItem
		err = data.FindOne(newsItemCollection, bson.M{"sourceurl": extractedSourceURL}, &existing)

		if err != nil && err != mongo.ErrNoDocuments {
			glog.Errorf("error when checking if the item has already been loaded. error: %s\n", err.Error())
			continue
		}

		if existing.SourceURL != "" {
			glog.Infoln("item already found, not adding this to the store...")
			continue
		}

		glog.Infoln("item not found, adding this to the store...")
		var result *mongo.InsertOneResult
		result, err = data.InsertOne(newsItemCollection, newsItem)
		if err != nil {
			glog.Errorf("adding item to the store failed with error: %s. Skipping this item \n", err.Error())
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
