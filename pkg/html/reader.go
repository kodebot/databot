package html

import (
	"io/ioutil"
	"net/http"

	"github.com/kodebot/databot/pkg/databot"

	"github.com/kodebot/databot/pkg/logger"
)

// DocumentReader is the abstract html document reader
type DocumentReader interface {
	ReadAsString() (string, error)
}

type documentReader struct {
	url string
}

type cachedDocumentReader struct {
	*documentReader
	databot.CacheService
}

// NewDocumentReader returns a new document reader to read html document from the given Url
func NewDocumentReader(url string) DocumentReader {
	return documentReader{url}
}

// NewCachedDocumentReader returns a new document reader of html document that is backed by given cache
func NewCachedDocumentReader(url string, cacheService databot.CacheService) DocumentReader {
	docReader := documentReader{url}
	return cachedDocumentReader{&docReader, cacheService}
}

func (d cachedDocumentReader) ReadAsString() (string, error) {
	cached := d.Get(d.url)
	if cached != nil {
		return cached.(string), nil
	}

	hot, err := d.ReadAsString()
	if err != nil {
		return "", err
	}

	d.Add(d.url, hot)
	return hot, nil
}

// ReadAsString returns http response as a string for given url
func (d documentReader) ReadAsString() (string, error) {
	var client http.Client
	resp, err := client.Get(d.url)
	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Errorf("error when retrieving raw feed from url %s status code: %d. error: %s\n", d.url, resp.StatusCode, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("error when reading body from url %s. error: %s\n", d.url, err.Error())
		return "", err
	}
	bodyStr := string(bodyBytes)
	return bodyStr, nil
}
