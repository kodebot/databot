package services

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

func getFeedConfig(t *testing.T, filePath string) models.FeedConfig {
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile(filePath, &feedConfig)
	if err != nil {
		t.Errorf("error when loading feed config: %s\n", err.Error())
		t.FailNow()
	}
	return feedConfig
}

func TestParseFeedWithDinamalarSportsBadmintonFeed(t *testing.T) {
	feedConfig := getFeedConfig(t, "./test_feed_configs/dinamalar_sports_badminton.toml")

	for _, feed := range feedConfig.Feed {
		t.Logf("processing %s \n", feed.URL)
		result := ParseFeed(feed)
		if result == nil {
			t.Errorf("failed when parsing feed from %s", feed.URL)
		}

		for _, item := range result {

			if item.Published == "" {
				t.Errorf("date not found for %s", item.Link)
				t.Fail()
			}

			if item.PublishedParsed == nil {
				t.Errorf("parsed date not found for %s", item.Link)
				t.Fail()
			}
		}

		t.Logf("finished processing %s \n", feed.URL)
	}
}

func TestCreateArticlesWithDinamalarSportsBadmintonFeed(t *testing.T) {
	feedConfig := getFeedConfig(t, "./test_feed_configs/dinamalar_sports_badminton.toml")

	for _, feed := range feedConfig.Feed {
		t.Logf("processing %s \n", feed.URL)
		result := ParseFeed(feed)
		if result == nil {
			t.Errorf("failed when parsing feed from %s", feed.URL)
		}

		articles := CreateArticles(result, feed)
		t.Logf("number of articles %d\n", len(articles))

		for _, article := range articles {
			t.Errorf("published date %s\n", article.PublishedDate)
		}
		t.Fail()
	}
}

func TestCreateArticlesWithDinamalarCinemaFeed(t *testing.T) {
	feedConfig := getFeedConfig(t, "./test_feed_configs/dinamalar_cinema.toml")

	for _, feed := range feedConfig.Feed {
		t.Logf("processing %s \n", feed.URL)
		result := ParseFeed(feed)
		if result == nil {
			t.Errorf("failed when parsing feed from %s", feed.URL)
		}

		articles := CreateArticles(result, feed)
		t.Logf("number of articles %d\n", len(articles))

		for _, article := range articles {
			t.Errorf("published date %s\n", article.PublishedDate)
		}
		t.Fail()
	}
}
