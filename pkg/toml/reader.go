package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
)

// FeedSpecReader represents toml feed config reader
type FeedSpecReader struct {
}

// Get creates Feed from the given toml file
func (reader *FeedSpecReader) Read(filePath string) databot.FeedSpec {

	var feed databot.FeedSpec
	_, err := toml.DecodeFile(filePath, &feed)
	if err != nil {
		logger.Errorf("error when loading feed info: %s\n", err.Error())
	}
	return feed
}

// FromString creates Feed from the given toml string
func (reader *FeedSpecReader) FromString(specContent string) databot.FeedSpec {

	var feed databot.FeedSpec
	_, err := toml.Decode(specContent, &feed)
	if err != nil {
		logger.Errorf("error when loading feed info: %s\n", err.Error())
	}
	return feed
}
