package services

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

type testFunc func(t *testing.T)

func test(t *testing.T, name string, test testFunc) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Helper()
		test(t)
		if t.Failed() {
			t.Errorf(">>>%s FAILED\n\n", name)
			return
		}
		t.Logf("%s SUCCEEDED\n\n", name)
	})
}

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

	// todo: test fixing illegal chars

	testArticles := CreateArticles(parsedFeeds, *feedConfig)

	testArticle := testArticles[0]

	// Title tests
	test(t, "Title should be same without leading and tailing whitespaces", func(t *testing.T) {
		expected := "கட்சி தாவுகிறாரா திவ்யா ஸ்பந்தனா"
		actual := testArticle.Title
		if actual != expected {
			t.Errorf("expect Title to be **%s** but found **%s**", expected, actual)
		}
	})

	// SourceUrl tests
	test(t, "SourceURL should match the link in feed", func(t *testing.T) {
		expected := "http://www.dinamalar.com/news_detail.asp?id=2289413"
		actual := testArticle.SourceURL
		if actual != expected {
			t.Errorf("expect Title to be **%s** but found **%s**", expected, actual)
		}
	})

	// ShortContent tests
	test(t, "ShortContent should be empty", func(t *testing.T) {
		expected := ""
		actual := testArticle.ShortContent
		if actual != expected {
			t.Errorf("expect ShortContent to be **%s** but found **%s**", expected, actual)
		}
	})

	test(t, "ShortContent should match feed when available", func(t *testing.T) {
		// title should be as is
		t.Skipf("todo test data not available\n\n")
	})

	// PublishedDate
	test(t, "PublishedDate should be UTC equivalent of feed date", func(t *testing.T) {
		expected := time.Date(2019, 6, 2, 16, 3, 0, 0, time.UTC) // Sun, 02 Jun 2019 21:33:00 +0530
		actual := testArticle.PublishedDate
		if !actual.Equal(expected) {
			t.Errorf("expect PublishedDate to be **%s** but found **%s**", expected, actual)
		}
	})

	/*
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
	*/

}
