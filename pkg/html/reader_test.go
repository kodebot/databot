package html

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/stretchr/testify/assert"
)

type testHTTPClient struct {
	getFn func(string) (*http.Response, error)
}

func (c *testHTTPClient) Get(url string) (*http.Response, error) {
	return c.getFn(url)
}

func newTestDocumentReader(url string, client *testHTTPClient) DocumentReader {
	return &documentReader{client, url}
}

func newTestCachedDocumentReader(url string, client *testHTTPClient, cacheManager *cache.Manager) DocumentReader {
	docReader := documentReader{client, url}
	return &cachedDocumentReader{&docReader, *cacheManager}
}

func TestReadAsStringWithoutCache(t *testing.T) {

	tests := []struct {
		httpError  bool
		statusCode int
		body       interface{}
		expected   string
	}{
		{true, 200, "bar", ""},
		{false, 200, "bar", "bar"},
		{false, 401, "bar", ""},
		{false, 200, nil, ""},
	}

	for _, test := range tests {
		getFn := func(url string) (*http.Response, error) {
			if test.httpError {
				return nil, errors.New("foo")
			}

			var body io.ReadCloser
			if test.body == nil {
				body = ioutil.NopCloser(bytes.NewReader([]byte(nil)))
			} else {
				body = ioutil.NopCloser(bytes.NewReader([]byte(test.body.(string))))
			}

			return &http.Response{StatusCode: test.statusCode, Body: body}, nil
		}

		reader := newTestDocumentReader("foo", &testHTTPClient{getFn})
		actual, _ := reader.ReadAsString()
		assert.Equal(t, test.expected, actual)
	}
}

func TestReadAsStringWithCache(t *testing.T) {

	tests := []struct {
		name       string
		cached     bool
		httpError  bool
		statusCode int
		expected   string
	}{
		{"not in cache, http error", false, true, 200, ""},
		{"not in cache, http 200", false, false, 200, "hot"},
		{"not in cache, http not 200", false, false, 401, ""},
		{"in cache, http error", true, true, 200, "cold"},
		{"in cache, http 200", true, false, 200, "cold"},
		{"in cache", true, false, 401, "cold"},
	}

	for _, test := range tests {
		getFn := func(url string) (*http.Response, error) {
			if test.httpError {
				return nil, errors.New("foo")
			}
			return &http.Response{StatusCode: test.statusCode, Body: ioutil.NopCloser(bytes.NewReader([]byte("hot")))}, nil
		}

		url := "foo"
		cache := cache.NewMemCache()

		if test.cached {
			cache.Add(url, "cold")
		}

		reader := newTestCachedDocumentReader(url, &testHTTPClient{getFn}, &cache)
		actual, err := reader.ReadAsString()
		assert.Equal(t, test.expected, actual, test.name)

		if !test.cached && err == nil {
			assert.Equal(t, actual, cache.Get(url), "should add to cache", test.name)
		}

		if !test.cached && err != nil {
			assert.Equal(t, nil, cache.Get(url), "should not be in cache", test.name)
		}
	}
}
