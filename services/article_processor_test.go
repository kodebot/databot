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
			if item.PublishedParsed == nil {
				t.Errorf("date not found for %s", item.Link)
			}
		}

		t.Logf("finished processing %s \n", feed.URL)
	}
}
