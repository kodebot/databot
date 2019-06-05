package services

import (
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

func loadFeedConfig(t *testing.T, filePath string) models.FeedConfig {
	if filePath == "" {
		filePath = "../conf/feed_parsing_config.toml"
	}
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile(filePath, &feedConfig)
	if err != nil {
		t.Fatalf("error when loading feed config: %s\n", err.Error())
	}
	return feedConfig
}

func getFeedConfigItemByName(t *testing.T, feedConfigName string, filePath string) *models.FeedConfigItem {
	feedConfigs := loadFeedConfig(t, filePath)

	for _, feedConfigItem := range feedConfigs.Feed {
		if feedConfigItem.Name == feedConfigName {
			return &feedConfigItem
		}
	}

	t.Fatalf("Feed config item with name %s not found", feedConfigName)
	return nil
}

func TestCreateArticle_dinamalar_politics(t *testing.T) {

	/*
	       name = "dinamalar_politics"
	       url = "http://rss.dinamalar.com/?cat=ara1"
	       origin = "dinamalar"
	       defaultCategory = "politics"

	       [feed.itemThumbImageExtractor]
	           sourceField = "Description"
	           scrapingRequired = true
	           selectorType = "regexp"
	           selector = "<img[^>]+src='(?P<Data>[^']+)"

	       [feed.itemUrlExtractor]
	           sourceField = "Link"
	           scrapingRequired = false
	           # selectorType = "regexp"
	           # selector = "<a[^>]+href='(?P<Data>[^']+)"

	   		General:
	   		1. check number of articles matches expected

	   		Fields:
	   			Title         string
	   			ShortContent  string
	   			PublishedDate time.Time
	   			Categories    []string
	   			ThumbImageURL string
	   			SourceURL     string
	   			Source        string
	   			OriginalFeed  gofeed.Item
	   			CreatedAt     time.Time


	*/
	feedConfigName := "dinamalar_politics"
	feedConfig := getFeedConfigItemByName(t, feedConfigName, "")

	if feedConfig == nil {
		t.Fatalf("Unable to load feed config with name %s", feedConfigName)
	}

	testFeedFile := "./test_data/feeds/dinamalar_politics_20190603.xml"
	testFeedXMLBytes, err := ioutil.ReadFile(testFeedFile)

	if err != nil {
		t.Fatalf("Unable to load test data from %s, %s", testFeedFile, err.Error())
	}

	parsedFeeds := ParseFeed(*feedConfig, string(testFeedXMLBytes))

	if parsedFeeds == nil {
		t.Fatalf("Unable to parse feed data from file %s with feed config name %s", testFeedFile, feedConfigName)
	}

	// General tests

	func() {
		// feed item count
		expected := 1
		actual := len(parsedFeeds)
		if actual != expected {
			t.Fatalf("expected **%d** parsed feeds but found **%d**", expected, actual)
		}
	}()

	testFeedItem := parsedFeeds[0]
	// Field level tests

	// Title tests
	func() {
		// title should be trimmed
		expected := "கட்சி தாவுகிறாரா திவ்யா ஸ்பந்தனா"
		actual := testFeedItem.Title
		if actual != expected {
			t.Fatalf("expect Title to be **%s** but found **%s**", expected, actual)
		}
	}()

}
