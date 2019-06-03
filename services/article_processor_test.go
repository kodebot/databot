package services

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

func getFeedConfig(t *testing.T, filePath string) models.FeedConfig {
	if filePath == "" {
		filePath = "../conf/feed_parsing_config.toml"
	}
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile(filePath, &feedConfig)
	if err != nil {
		t.Errorf("error when loading feed config: %s\n", err.Error())
		t.FailNow()
	}
	return feedConfig
}

func TestCreateArticle_dinamalar_politics(t *testing.T) {

	getFeedConfig()
	getRawFeedAsString()

}
