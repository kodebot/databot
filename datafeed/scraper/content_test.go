package scraper

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractConent(t *testing.T) {
	content := extractContent("https://www.dinamalar.com/news_detail.asp?id=2297141", "url", "#mvp-content-wrap")
	ioutil.WriteFile("./temp.html", []byte(content), os.ModePerm)
}
