package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/logger"
	"github.com/kodebot/newsfeed/models"
)

// AppSettings contain all the settings
var AppSettings models.AppSettings

// Init the values
func Init() {
	_, err := toml.DecodeFile("./conf/article_category_config.toml", &AppSettings)
	if err != nil {
		logger.Fatalf("error when loading article category config: %s\n", err.Error())
	}

	_, err = toml.DecodeFile("./conf/article_source_config.toml", &AppSettings)
	if err != nil {
		logger.Fatalf("error when loading article source config: %s\n", err.Error())
	}
}
