// +build integration

package datafeed

import (
	"testing"
	"time"

	"github.com/kodebot/newsfeed/datafeed/record"
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
)

func TestIntegrationParseDinamalarPoliticsRssFeed(t *testing.T) {

	url := "http://rss.dinamalar.com/?cat=ara1"

	recordSetting := record.Info{
		Fields: []field.Info{
			{
				Name: "Title",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Value}},
			{
				Name: "Description",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Value}},
			{
				Name: "Content",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Value}},
			{
				Name: "Published",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Value}},
			{
				Name: "PublishedDate",
				CollectorSetting: fcollectors.CollectorInfo{
					Type:       fcollectors.Value,
					Parameters: map[string]interface{}{"source": "PublishedParsed"}}},
			{
				Name: "ThumbImageUrl",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Regexp,
					Parameters: map[string]interface{}{
						"expr":   "<img[^>]+src='(?P<data>[^']+)",
						"source": "Description"}}},
			{
				Name: "SourceUrl",
				CollectorSetting: fcollectors.CollectorInfo{
					Type: fcollectors.Value,
					Parameters: map[string]interface{}{
						"source": "Link"}}},
		}}

	records := ParseFromURL(url, rcollectors.RssAtom, recordSetting)

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
		t.Fatalf("unable to retrieve published")
	} else {
		t.Logf("published retrieved successfully, Published: %s", published)
	}

	if publishedDate, ok := first["PublishedDate"].(*time.Time); !ok {
		t.Fatalf("unable to retrieve published date ")
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
