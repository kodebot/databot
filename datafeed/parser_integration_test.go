// build integration

package datafeed

import (
	"testing"
	"time"

	"github.com/kodebot/newsfeed/datafeed/collectors/record/fields"
	"github.com/kodebot/newsfeed/datafeed/model"
)

func TestIntegrationParseDinamalarPoliticsRssFeed(t *testing.T) {

	url := "http://rss.dinamalar.com/?cat=ara1"

	recordSetting := model.RecordInfo{
		Fields: []model.FieldInfo{
			{
				Name: "Title",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Value}},
			{
				Name: "Description",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Value}},
			{
				Name: "Content",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Value}},
			{
				Name: "Published",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Value}},
			{
				Name: "PublishedDate",
				CollectorSetting: fields.CollectorInfo{
					Type:       fields.Value,
					Parameters: map[string]interface{}{"source": "PublishedParsed"}}},
			{
				Name: "ThumbImageUrl",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Regexp,
					Parameters: map[string]interface{}{
						"expr":   "<img[^>]+src='(?P<data>[^']+)",
						"source": "Description"}}},
			{
				Name: "SourceUrl",
				CollectorSetting: fields.CollectorInfo{
					Type: fields.Value,
					Parameters: map[string]interface{}{
						"source": "Link"}}},
		}}

	records := ParseFromURL(url, model.RssAtom, recordSetting)

	if count := len(records); count < 1 {
		t.Fatalf("should have at least one record but found %d", count)
	}

	first := records[0]

	for key, val := range first {
		t.Logf("%s:%#v", key, val)
	}

	if title, ok := first["Title"].(string); !ok {
		t.Fatalf("unable to retrieve title")
	} else {
		t.Logf("title retrieved successfully, Title: %s", title)
	}

	if description, ok := first["Description"].(string); !ok {
		t.Fatalf("unable to retrieve description")
	} else {
		t.Logf("descritpion retrieved successfully, Description: %s", description)
	}

	if content, ok := first["Content"].(string); !ok {
		t.Fatalf("unable to retrieve content")
	} else {
		t.Logf("content retrieved successfully, Content: %s", content)
	}

	if published, ok := first["Published"].(string); !ok {
		t.Fatalf("unable to retrieve publihsed")
	} else {
		t.Logf("published retrieved successfully, Published: %s", published)
	}

	if publishedDate, ok := first["PublishedDate"].(*time.Time); !ok {
		t.Fatalf("unable to retrieve publihsed date ")
	} else {
		t.Logf("published date retrieved successfully, PublishedDate: %s", publishedDate)
	}

	if thumbImageURL, ok := first["ThumbImageUrl"].(string); !ok {
		t.Fatalf("unable to retrieve thumbImageUrl")
	} else {
		t.Logf("thumbImageUrl retrieved successfully, ThumbImageUrl: %s", thumbImageURL)
	}

	if sourceURL, ok := first["SourceUrl"].(string); !ok {
		t.Fatalf("unable to retrieve sourceUrl")
	} else {
		t.Logf("sourceUrl retrieved successfully, SourceUrl: %s", sourceURL)
	}
}
