// +build integration

package datafeed

import (
	"testing"

	cmodel "github.com/kodebot/newsfeed/datafeed/collectors/model"
	"github.com/kodebot/newsfeed/datafeed/model"
)

func TestIntegrationParseDinamalarPoliticsRssFeed(t *testing.T) {

	url := "http://rss.dinamalar.com/?cat=ara1"

	recordSetting := model.RecordSetting{
		FieldSettings: []model.FieldSetting{
			{
				Name: "Title",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Value}},
			{
				Name: "Description",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Value}},
			{
				Name: "Content",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Value}},
			{
				Name: "Published",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Value}},
			{
				Name: "ThumbImageUrl",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Regexp,
					Parameters: map[string]interface{}{
						"expr":   "<img[^>]+src='(?P<data>[^']+)",
						"source": "Description"}}},
			{
				Name: "SourceUrl",
				CollectorSetting: cmodel.FieldCollectorSetting{
					Type: cmodel.Value,
					Parameters: map[string]interface{}{
						"source": "Link"}}},
		}}

	records := ParseFromURL(url, model.RssAtom, recordSetting)

	if count := len(records); count < 1 {
		t.Fatalf("should have at least one record but found %d", count)
	}

	first := records[0]

	for key, val := range first {
		t.Logf("%s:%#v", key, *val)
	}

	if *first["Title"] == nil {
		t.Fatalf("unable to retrieve title")
	}

	if *first["Description"] == nil {
		t.Fatalf("unable to retrieve description")
	}

	if *first["Content"] == nil {
		t.Fatalf("unable to retrieve content")
	}

	if *first["Published"] == nil {
		t.Fatalf("unable to retrieve published date")
	}

	if *first["ThumbImageUrl"] == nil {
		t.Fatalf("unable to retrieve thumb image url")
	}

	if *first["SourceUrl"] == nil {
		t.Fatalf("unable to retrieve source url")
	}
}
