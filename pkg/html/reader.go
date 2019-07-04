package html

import (
	"io/ioutil"
	"net/http"

	"github.com/kodebot/databot/pkg/cache"

	"github.com/kodebot/databot/pkg/logger"
)

type httpClient interface {
	Get(url string) (*http.Response, error)
}

type defaultHTTPClient struct {
	client http.Client
}

func (c *defaultHTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// DocumentReader is the abstract html document reader
type DocumentReader interface {
	ReadAsString() (string, error)
}

type documentReader struct {
	client httpClient
	url    string
}

type cachedDocumentReader struct {
	*documentReader
	cache *cache.Manager
}

// NewDocumentReader returns a new document reader to read html document from the given Url
func NewDocumentReader(url string) DocumentReader {
	return &documentReader{&defaultHTTPClient{}, url}
}

// NewCachedDocumentReader returns a new document reader of html document that is backed by given cache
func NewCachedDocumentReader(url string, cacheManager *cache.Manager) DocumentReader {
	docReader := documentReader{&defaultHTTPClient{}, url}
	return &cachedDocumentReader{&docReader, cacheManager}
}

func (d *cachedDocumentReader) ReadAsString() (string, error) {
	URL := d.url
	cache := *(d.cache)
	cached := cache.Get(URL)
	if cached != nil {
		return cached.(string), nil
	}

	hot, err := d.documentReader.ReadAsString()
	if err != nil {
		return "", err
	}

	cache.Add(URL, hot)
	return hot, nil
}

// ReadAsString returns http response as a string for given url
func (d *documentReader) ReadAsString() (string, error) {
	resp, err := d.client.Get(d.url)
	if err != nil {
		logger.Errorf("error when retrieving raw feed from url %s. error: %s", d.url, err.Error())
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("error when retrieving raw feed from url %s status code: %d.", d.url, resp.StatusCode)
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
