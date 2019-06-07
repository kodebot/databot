package datafeed

import (
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/model"
)

// readAsString returns http response as a string for given url
func readAsString(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		glog.Errorf("error when retrieving raw feed from url %s status code: %d. error: %s\n", url, resp.StatusCode, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("error when reading body from url %s. error: %s\n", url, err.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func readRecordSettings(filePath string) []model.RecordSetting {

	var recordSettings struct {
		Record []model.RecordSetting
	}
	_, err := toml.DecodeFile(filePath, &recordSettings)
	if err != nil {
		glog.Errorf("error when loading feed config: %s\n", err.Error())
	}
	return recordSettings.Record
}
