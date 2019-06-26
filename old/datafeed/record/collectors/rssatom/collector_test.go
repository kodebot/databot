package rssatom

import (
	"testing"

	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"

	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
	ftransformers "github.com/kodebot/newsfeed/datafeed/record/collectors/field/transformers"
)

var testFeedDataWithSingleItem = `<?xml
version='1.0' encoding='utf-8'?>
<rss version='2.0'>
	<channel>
		<title>Dinamalar.com |ஜூன் 05,2019</title>
		<link>http://www.dinamalar.com</link>
		<managingEditor>coordinator@dinamalar.com(Editor)</managingEditor>
		<image>
			<title>dinamalar.com</title>
			<url>http://img.dinamalar.com/images/top.png</url>
			<link>http://www.dinamalar.com</link>
			<width>150</width>
			<height>40</height>
			<description>Visit dinamalar.com</description>
		</image>
		<item>
			<title>  பறவைகளுக்கு விலாசம் சொன்னது யார்?  </title>
			<link>http://www.dinamalar.com/news_detail.asp?id=2290872</link>
			<category></category>
			<language>ta</language>
			<pubDate>Tue, 04 Jun 2019 23:34:00 +0530</pubDate>
			<description><![CDATA[<a href='http://www.dinamalar.com/news_detail.asp?id=2290872'><img height='65' width='65' src='http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2290872_150_100.jpg' border='0' style='margin: 10px 0 0 10px; padding: 4px; ' align='left'  /></a><p>...</p>]]></description>
		</item>     
	</channel>
</rss>`

func TestCollectRssAtomFieldCollectorWithFieldName(t *testing.T) {

	fieldsInfo := []field.Info{{
		Name: "Title",
		CollectorInfo: fcollectors.CollectorInfo{
			Type: fcollectors.RssAtomField}}}

	result := Collect(testFeedDataWithSingleItem, fieldsInfo)

	if count := len(result); count != 1 {
		t.Fatalf("expected 1 collected record but found %d", count)
	}

	expectedTitle := "பறவைகளுக்கு விலாசம் சொன்னது யார்?"
	if title := result[0]["Title"]; title != expectedTitle {
		t.Fatalf("collected value doesn't match the expected. Expected: %s ** Actual: %s", expectedTitle, title)
	}
}

func TestCollectRssAtomFieldCollectorWithSourceParameter(t *testing.T) {

	fieldsInfo := []field.Info{{
		Name: "Header",
		CollectorInfo: fcollectors.CollectorInfo{
			Type: fcollectors.RssAtomField,
			Parameters: map[string]interface{}{
				"source": "Title"}}}}

	result := Collect(testFeedDataWithSingleItem, fieldsInfo)

	if count := len(result); count != 1 {
		t.Fatalf("expected 1 collected record but found %d", count)
	}

	expectedValue := "பறவைகளுக்கு விலாசம் சொன்னது யார்?"
	if value := result[0]["Header"]; value != expectedValue {
		t.Fatalf("collected value doesn't match the expected. Expected: %s ** Actual: %s", expectedValue, value)
	}
}

func TestCollectRssAtomFieldCollectorWithTransformer(t *testing.T) {

	fieldsInfo := []field.Info{{
		Name: "Title",
		CollectorInfo: fcollectors.CollectorInfo{
			Type: fcollectors.RssAtomField},
		TransformersInfo: []ftransformers.TransformerInfo{{
			Transformer: ftransformers.Empty}}}}

	result := Collect(testFeedDataWithSingleItem, fieldsInfo)

	if count := len(result); count != 1 {
		t.Fatalf("expected 1 collected record but found %d", count)
	}

	if actualValue := (result[0])["Title"]; actualValue != "" {
		t.Fatalf("collected value doesn't match the expected. Expected: <<EMPTY>> ** Actual: %s", actualValue)
	}
}

func TestCollectUnknownCollector(t *testing.T) {

	fieldsInfo := []field.Info{{
		Name: "Description",
		CollectorInfo: fcollectors.CollectorInfo{
			Type: fcollectors.Unknown}}}

	result := Collect(testFeedDataWithSingleItem, fieldsInfo)

	if count := len(result); count != 1 {
		t.Fatalf("expected 1 collected record but found %d", count)
	}

	if actualValue := (result[0])["Description"]; actualValue != nil {
		t.Fatalf("collected value doesn't match the expected. Expected: NIL ** Actual: %s", actualValue)
	}
}
