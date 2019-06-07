package feed

import (
	"testing"

	"github.com/kodebot/newsfeed/datafeed/collectors/model"
)

var testFeedDataWithSingleItem string = `<?xml
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

func TestCollect_single_VALUE_collector(t *testing.T) {

	fieldCollectors := []model.FieldCollectorSetting{{
		Field: "Title",
		Type:  model.VALUE}}

	result := Collect(testFeedDataWithSingleItem, fieldCollectors)

	if count := len(result); count != 1 {
		t.Fatalf("expected 1 collected record but found %d", count)
	}

	expectedTitle := "பறவைகளுக்கு விலாசம் சொன்னது யார்?"
	if title := *(result[0])["Title"]; title != expectedTitle {
		t.Fatalf("collected value doesn't match the expected. Expected: %s ** Actual: %s", expectedTitle, title)
	}

}
