package services

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
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

// ParseFeed from the given url
func ParseFeed(feedConfig models.FeedConfigItem) []*gofeed.Item {
	defer handleLoadFeedPanics(feedConfig)
	glog.Infof("starting loading articles from URL: %s \n", feedConfig.URL)
	defer glog.Infof("ending loading feed from URL: %s", feedConfig.URL)

	xmlString, err := getRawFeedAsString(feedConfig.URL)

	if err != nil {
		glog.Errorf("retrieving feed xml failed with error %s. Skipping this source.\n", err.Error())
		return nil
	}

	parser := gofeed.NewParser()
	feed, err := parser.ParseString(xmlString)
	if err != nil {
		glog.Errorf("parsing feed failed with error %s. Skipping this source.\n", err.Error())
		dumpErrorInDatabase("parsing feed failed", errorDump{
			FeedConfig: feedConfig,
			Error:      err.Error()})
		return nil
	}

	totalItems := len(feed.Items)
	if totalItems == 0 {
		glog.Infoln("no items found to process")
	} else {

		glog.Infof("%d items found\n", totalItems)
	}

	return feed.Items

}

// CreateArticles loads news articles into the database
func CreateArticles(feedItems []*gofeed.Item, feedConfig models.FeedConfigItem) []*models.Article {
	glog.Info("creating articles...")
	result := []*models.Article{}
	totalItems := len(feedItems)

	for i, item := range feedItems {
		glog.Infof("creating article %d of %d\n", i+1, totalItems)

		glog.Infof("item data dump %+v\n", item)

		var revisedDate *time.Time
		if item.PublishedParsed != nil {
			revisedDate = item.PublishedParsed
		} else {
			glog.Infoln("parsed date not found. trying extractors...")
			for _, extractor := range feedConfig.PublishedDateExtractors {
				extractedDate, err := extractData(item, extractor)
				if err != nil {
					glog.Warningf("unable to extract date using date extractor: %s\n", err.Error())
					continue
				}
				dateString, err := convertTamilToEnglishDate(extractedDate)

				if err != nil {
					glog.Warningf("unable to convert extracted tamil date %s to english: %s\n", extractedDate, err.Error())
					continue
				}

				var parsedExtractedDate time.Time

				for _, layout := range feedConfig.DateLayouts {
					var err error
					parsedExtractedDate, err = time.Parse(layout, dateString)
					if err != nil {
						glog.Warningf("unable to parse the date extracted using layout %s error: %s\n", layout, err.Error())
						continue
					}
					break
				}

				revisedDate = &parsedExtractedDate
				glog.Infof("extracted date %s\n", revisedDate)
			}
		}

		if revisedDate == nil && feedConfig.UseCurrentDateTimeWhenPubDateMissing == true {
			glog.Infoln("published date missing time, setting it current date time...")
			x := time.Now()
			revisedDate = &x
		}

		if revisedDate == nil {
			glog.Errorf("unable to establish date for the article at all. Skipping this item.\n")
			dumpErrorInDatabase("extracting source link failed", errorDump{
				FeedConfig: feedConfig,
				OtherData:  item})
			continue
		}

		if isUtcMidnight(revisedDate) { // this means we don't have time - only date is present - just use current time
			glog.Infoln("published date missing time, setting it current time...")
			now := time.Now()
			if revisedDate.Before(now) {
				// when processing feed from india late in the evening the feed will come for next day
				// adding current UK time to next day feed from India result in next day + current UK time
				// if this happens the feed will have future published date - fall back to time.Now() for this cases
				x := revisedDate.Add(time.Hour*time.Duration(now.Hour()) +
					time.Minute*time.Duration(now.Minute()) +
					time.Second*time.Duration(now.Second()))
				revisedDate = &x
			} else {
				revisedDate = &now
			}
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

		extractedShortContent := ""
		if feedConfig.ShortContentExtractor != (models.FeedDataExtractorConfig{}) {
			glog.Infoln("extracting short content...")
			var err error
			extractedShortContent, err = extractData(item, feedConfig.ShortContentExtractor)

			if err != nil {
				glog.Warningf("extracting short content failed with error: %s\n", err.Error())
			} else if extractedShortContent == "" {
				glog.Warningln("extracting short content failed - no content found.")
			} else {
				glog.Infof("extracted short content successfully")
			}
		}

		article := models.Article{
			Title:         item.Title,
			ShortContent:  extractedShortContent,
			PublishedDate: *revisedDate,
			Categories:    []string{feedConfig.DefaultCategory},
			ThumbImageURL: extactedImageURL,
			SourceURL:     extractedSourceURL,
			Source:        feedConfig.Origin,
			OriginalFeed:  *item,
			CreatedAt:     time.Now()}

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

// some rss feed has illegal xml data - this is to replace them with empty letter
func getRawFeedAsString(url string) (string, error) {
	illegalXMLCharacters := []rune{
		'\u0001', '\u0002', '\u0003', '\u0004', '\u0005', '\u0006', '\u0007',
		'\u0008', '\u000b', '\u000c', '\u000e', '\u000f', '\u0010', '\u0011',
		'\u0012', '\u0013', '\u0014', '\u0015', '\u0016', '\u0017', '\u0018',
		'\u0019', '\u001a', '\u001b', '\u001c', '\u001d', '\u001e', '\u001f'}

	var client http.Client
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		glog.Errorf("error when retrieving raw feed from url %s status code: %d. error: %s\n", url, resp.StatusCode, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("error when reading body from url %s. error: %s\n", url, err.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	correctedBodyString := bodyString

	for _, char := range illegalXMLCharacters {
		correctedBodyString = strings.Replace(correctedBodyString, string(char), "", -1)
	}

	return correctedBodyString, nil
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
			if val == "Data" {
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
