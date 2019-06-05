package services

import (
	"fmt"
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
	actualArticle := testArticles[0]
	expectedArticle := models.Article{
		Title:         "கட்சி தாவுகிறாரா திவ்யா ஸ்பந்தனா",
		ShortContent:  "",
		PublishedDate: time.Date(2019, 6, 2, 16, 3, 0, 0, time.UTC), // Sun, 02 Jun 2019 21:33:00 +0530
		SourceURL:     "http://www.dinamalar.com/news_detail.asp?id=2289413",
		Categories:    []string{"politics"},
		ThumbImageURL: "http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2289413_150_100.jpg",
		Source:        "dinamalar"}

	test(t, fmt.Sprintf("Dinamalar politics feed %s", testFeedFile), func(t *testing.T) {
		if actualArticle.Title != expectedArticle.Title {
			t.Errorf("Title >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.Title, actualArticle.Title)
		}

		if actualArticle.ShortContent != expectedArticle.ShortContent {
			t.Errorf("ShortContent >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.ShortContent, actualArticle.ShortContent)
		}

		if !actualArticle.PublishedDate.Equal(expectedArticle.PublishedDate) {
			t.Errorf("PublishedDate >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.PublishedDate, actualArticle.PublishedDate)
		}

		if len(actualArticle.Categories) != len(expectedArticle.Categories) {
			t.Errorf("Categories count >> EXPECTED: **%d** ACTUAL: **%d**", len(expectedArticle.Categories), len(actualArticle.Categories))
		}

		for _, expectedCategory := range expectedArticle.Categories {
			found := false
			for _, actualCategory := range actualArticle.Categories {
				if expectedCategory == actualCategory {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Category missing >> EXPECTED: **%s** but not found", expectedCategory)
			}
		}

		if actualArticle.ThumbImageURL != expectedArticle.ThumbImageURL {
			t.Errorf("ThumbImageURL >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.ThumbImageURL, actualArticle.ThumbImageURL)
		}

		if actualArticle.SourceURL != expectedArticle.SourceURL {
			t.Errorf("SourceURL >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.SourceURL, actualArticle.SourceURL)
		}

		if actualArticle.Source != expectedArticle.Source {
			t.Errorf("Source >> EXPECTED: **%s** ACTUAL: **%s**", expectedArticle.Source, actualArticle.Source)
		}

		now := time.Now()
		twoSecBefore := now.Add(time.Second * -2)
		if !((actualArticle.CreatedAt.Equal(now) || actualArticle.CreatedAt.Before(now)) &&
			actualArticle.CreatedAt.After(twoSecBefore)) {
			t.Errorf("CreatedAt EXPECTED: between **%s** and **%s** ACTUAL: **%s**", twoSecBefore, now, actualArticle.CreatedAt)
		}
	})
}
