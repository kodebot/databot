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
		"https://cinema.dinamalar.com/tamil-news/78958/cinema/Kollywood/aruvam-is-social-triller-movie.htm",
		"url",
		[]string{"#selDetail"},
		[]string{"#selDetail > h2",
			"#selDetail > div.darrow.clsFloatleft",
			"#selDetail > div.date.clsFloatright",
			"#selDetail > div.bar_cont",
			"#selDetail > div:nth-child(4)",
			"#selDetail > div:nth-child(6)"})

	ioutil.WriteFile("./temp.html", []byte(content), os.ModePerm)
}
