package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
)

// FeedConfigReader represents toml feed config reader
type FeedConfigReader struct {
}

// Get creates Feed from the given toml file
func (reader *FeedConfigReader) Get(filePath string) databot.Feed {

	var feed databot.Feed
	_, err := toml.DecodeFile(filePath, &feed)
	if err != nil {
		logger.Errorf("error when loading feed info: %s\n", err.Error())
	}
	return feed
}
