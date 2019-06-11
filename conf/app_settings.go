package conf

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/logger"
	"github.com/kodebot/newsfeed/models"
)

// AppSettings contain all the settings
var AppSettings models.AppSettings

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("unable to get the current working dir")
	}

	path := strings.Split(pwd, "newsfeed")[0] + "newsfeed/conf/"

	_, err = toml.DecodeFile(path+"article_category_config.toml", &AppSettings)
	if err != nil {
		logger.Fatalf("error when loading article category config: %s\n", err.Error())
	}

	_, err = toml.DecodeFile(path+"article_source_config.toml", &AppSettings)
	if err != nil {
		logger.Fatalf("error when loading article source config: %s\n", err.Error())
	}

	appSettingsFile := path + "app_settings.toml"
	if os.Getenv("env") == "PROD" {
		appSettingsFile = path + "app_settings_PROD.toml"
	}

	_, err = toml.DecodeFile(appSettingsFile, &AppSettings)
	if err != nil {
		logger.Fatalf("error when loading app settings. error %s ", err.Error())
	}
}
