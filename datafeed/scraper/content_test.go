package scraper

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractConent(t *testing.T) {
	// content := extractContent(
	// 	"http://www.dinamalar.com/news_detail.asp?id=2296694",
	// 	"url",
	// 	[]string{"#mvp-content-wrap", "#columns > div:nth-child(8) > div.row > div.col-sm-12.col-md-8"})

	// content := extractContent(
	// 	"http://www.dinamalar.com/news_detail.asp?id=2296595",
	// 	"url",
	// 	[]string{"#mvp-content-wrap", "#columns > div:nth-child(8) > div.row > div.col-sm-12.col-md-8"})

	content := extractContent(
		"https://www.dinamalar.com/news_detail.asp?id=2297248",
		"url",
		[]string{"#mvp-content-wrap", "#columns > div:nth-child(8) > div.row > div.col-sm-12.col-md-8"})

	ioutil.WriteFile("./temp.html", []byte(content), os.ModePerm)
}
