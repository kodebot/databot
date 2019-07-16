package processor

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/stringutil"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("http:get", cssSelect)
}

func httpGet(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
	useCacheParam := params["useCache"]
	useCache := false

	if useCacheParam != nil {
		if _, ok := useCacheParam.(bool); ok {
			useCache = true
		}
	}

	output := make(chan interface{})
	go func() {
		for newInput := range input {
			url, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", url)
			}

			var docReader html.DocumentReader
			if useCache {
				docReader = html.NewCachedDocumentReader(url, cache.Current())
			} else {
				docReader = html.NewDocumentReader(url)
			}

			htmlStr, err := docReader.ReadAsString()
			if err != nil {
				logger.Errorf("unable to get html from url: %s, skipping it", url)
			} else {
				htmlStr = fixRelativePaths(url, htmlStr)
				output <- htmlStr
			}
		}
		close(output)
	}()

	return output
}

func fixRelativePaths(sourceURL string, htmlStr string) string {
	matches := stringutil.RegexpMatchAll(htmlStr, "(href|src)=(\"|')(?P<data>[^(\"|')]+)")

	for _, match := range matches {
		absolutePath := resolveRelativePath(sourceURL, match)
		replacer := strings.NewReplacer(fmt.Sprintf("\"%s\"", match), fmt.Sprintf("\"%s\"", absolutePath), fmt.Sprintf("'%s'", match), fmt.Sprintf("'%s'", absolutePath))
		htmlStr = replacer.Replace(htmlStr)
	}
	return htmlStr
}

func resolveRelativePath(sourceURL string, relativePath string) string {
	if strings.HasPrefix(relativePath, "http") {
		return relativePath
	} else if strings.HasPrefix(relativePath, "/") {
		baseURL := regexp.MustCompile("^.+?[^/:]([?/]|$)").FindString(sourceURL)
		baseURL = strings.TrimRight(baseURL, "/") // remove tailing / if present
		return baseURL + relativePath
	} else if strings.HasSuffix(sourceURL, "/") {
		return sourceURL + relativePath
	} else {
		return sourceURL + "/" + relativePath
	}
}
