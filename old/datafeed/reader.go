package datafeed

import (
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/logger"
)

// readAsString returns http response as a string for given url
func readAsString(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Errorf("error when retrieving raw feed from url %s status code: %d. error: %s\n", url, resp.StatusCode, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("error when reading body from url %s. error: %s\n", url, err.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func readFeedInfo(filePath string) FeedInfo {

	var feedInfo FeedInfo
	_, err := toml.DecodeFile(filePath, &feedInfo)
	if err != nil {
		logger.Errorf("error when loading feed info: %s\n", err.Error())
	}
	return feedInfo
}
