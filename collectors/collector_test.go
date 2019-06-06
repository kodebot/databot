package collectors

import (
	"fmt"
	"testing"

	"github.com/kodebot/newsfeed/collectors/model"
)

func TestCollect(t *testing.T) {

	var fieldSettings []model.FieldCollectorSetting
	fieldSettings = append(fieldSettings, model.FieldCollectorSetting{Field: "title", Type: model.VALUE})

	result := Collect("http://rss.vikatan.com/feeds/politics_news.rss", model.FEED, fieldSettings)

	for _, record := range result {
		t.Logf(string(fmt.Sprintf("%s", record["title"])))
	}
}
