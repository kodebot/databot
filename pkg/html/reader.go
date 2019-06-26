package html

import (
	"io/ioutil"
	"net/http"

	"github.com/kodebot/databot/pkg/logger"
)

// GetRespBodyAsStr returns http response as a string for given url
func GetRespBodyAsStr(url string) (string, error) {
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
	bodyStr := string(bodyBytes)
	return bodyStr, nil
}
