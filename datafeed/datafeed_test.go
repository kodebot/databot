package datafeed

import (
	"testing"

	"github.com/kodebot/newsfeed/datafeed/record"
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
	ftransformers "github.com/kodebot/newsfeed/datafeed/record/collectors/field/transformers"
)

func TestNewRssAtom(t *testing.T) {

	data := `<?xml
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
			<item>
            <title>  ஒட்டுமொத்த கல்வி முறையையும் சீரமையுங்கள்: உச்ச நீதிமன்றம் உத்தரவு</title>
            <link>http://www.dinamalar.com/news_detail.asp?id=2291013</link>
            <category></category>
            <language>ta</language>
            <pubDate>Wed, 05 Jun 2019 01:43:00 +0530</pubDate>
            <description><![CDATA[<a href='http://www.dinamalar.com/news_detail.asp?id=2291013'><img height='65' width='65' src='http://img.dinamalar.com/data/thumbnew/Tamil_News_thumb_2291013_150_100.jpg' border='0' style='margin: 10px 0 0 10px; padding: 4px; ' align='left'  /></a><p>...</p>]]></description>
        </item>   
		</channel>
	</rss>`

	var recordInfo record.Info

	recordInfo.Fields = []field.Info{
		{
			Name: "Title",
			CollectorInfo: fcollectors.CollectorInfo{
				Type: fcollectors.RssAtomField},
			TransformersInfo: []ftransformers.TransformerInfo{{
				Transformer: ftransformers.Trim}},
		}, {
			Name: "PublishedParsed",
			CollectorInfo: fcollectors.CollectorInfo{
				Type: fcollectors.RssAtomField},
			TransformersInfo: []ftransformers.TransformerInfo{{
				Transformer: ftransformers.FormatDate}}},
	}

	records := New(data, rcollectors.RssAtom, recordInfo)

	if count := len(records); count != 2 {
		t.Fatalf("expected 2 records but found %d", count)
	}

	expectedResults := []struct {
		Title           string
		PublishedParsed string
	}{
		{
			"பறவைகளுக்கு விலாசம் சொன்னது யார்?",
			"2019-06-04 18:04:00 +0000 UTC"},
		{
			"ஒட்டுமொத்த கல்வி முறையையும் சீரமையுங்கள்: உச்ச நீதிமன்றம் உத்தரவு",
			"2019-06-04 20:13:00 +0000 UTC"}}

	for i, expectedResult := range expectedResults {
		if title := records[i]["Title"]; title != expectedResult.Title {
			t.Fatalf("parsed item Title doesn't match the expected. Expected: %s ** Actual: %s", expectedResult.Title, title)
		}

		if publishedParsed := records[i]["PublishedParsed"]; (publishedParsed.(string)) != expectedResult.PublishedParsed {
			t.Fatalf("parsed item published doesn't match the expected.  Expected: %s ** Actual: %s", expectedResult.PublishedParsed, (publishedParsed.(string)))
		}
	}
}
