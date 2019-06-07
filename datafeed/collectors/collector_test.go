package collectors

import (
	"fmt"
	"testing"

	"github.com/kodebot/newsfeed/datafeed/collectors/model"
)

func TestCollect(t *testing.T) {

	var fieldSettings []model.FieldCollectorSetting
	fieldSettings = append(fieldSettings, model.FieldCollectorSetting{Field: "Title", Type: model.VALUE})
	fieldSettings = append(fieldSettings, model.FieldCollectorSetting{Field: "Description", Type: model.VALUE})
	fieldSettings = append(fieldSettings, model.FieldCollectorSetting{Field: "PublishedDate", Type: model.VALUE})

	imageParameters := make(map[string]interface{})
	imageParameters["Source"] = "Description"
	imageParameters["Expr"] = "<img[^>]+src=\"(?P<data>[^\"]+)"
	fieldSettings = append(fieldSettings, model.FieldCollectorSetting{Field: "ImageUrl", Type: model.REGEXP, Parameters: imageParameters})

	result := Collect("http://rss.vikatan.com/feeds/politics_news.rss", model.FEED, fieldSettings)

	for _, record := range result {
		t.Logf(string(fmt.Sprintf("%s", record["Title"])))
		t.Logf(string(fmt.Sprintf("%s", record["Description"])))
		t.Logf("%s", (record["PublishedDate"]))
		t.Logf("%s", (record["ImageUrl"]))
	}
}
