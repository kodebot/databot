package feed

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/mmcdole/gofeed"
)

type feedItem struct {
	Title               string
	Description         string
	ImageURL            string
	ImageTitle          string
	PublishedDateString string
	PublishedDate       time.Time
	Category            string
	Content             string
	SiteURL             string
	OriginalItem        interface{}
}

func readFromURL(URL string) []*feedItem {
	glog.Infof("reading feed from URL: %s \n", URL)
	defer glog.Infof("ending loading feed from URL: %s", URL)

	xmlString, err := getRawFeedAsString(URL)
	if err != nil {
		glog.Errorf("retrieving feed xml failed with error %s. Skipping this source.\n", err.Error())
		return nil
	}

	xmlString = fixIllegalXMLCharacters(xmlString)
	return readFromXML(xmlString)
}

func readFromXML(xmlString string) []*feedItem {
	glog.Infof("parsing feed xml")
	parser := gofeed.NewParser()
	feed, err := parser.ParseString(xmlString)
	if err != nil {
		glog.Errorf("parsing feed failed with error %s.", err.Error())
		return nil
	}

	totalItems := len(feed.Items)
	if totalItems == 0 {
		glog.Infoln("no items found")
	} else {

		glog.Infof("%d items found\n", totalItems)
	}

	var feedItems []*feedItem
	for _, item := range feed.Items {
		feedItems = append(feedItems, toFeedItem(item))
	}

	return feedItems
}

func getRawFeedAsString(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		glog.Errorf("error when retrieving raw feed from url %s status code: %d. error: %s\n", url, resp.StatusCode, err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("error when reading body from url %s. error: %s\n", url, err.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func fixIllegalXMLCharacters(xmlString string) string {
	illegalXMLCharacters := []rune{
		'\u0001', '\u0002', '\u0003', '\u0004', '\u0005', '\u0006', '\u0007',
		'\u0008', '\u000b', '\u000c', '\u000e', '\u000f', '\u0010', '\u0011',
		'\u0012', '\u0013', '\u0014', '\u0015', '\u0016', '\u0017', '\u0018',
		'\u0019', '\u001a', '\u001b', '\u001c', '\u001d', '\u001e', '\u001f'}

	correctedBodyString := xmlString

	for _, char := range illegalXMLCharacters {
		correctedBodyString = strings.Replace(correctedBodyString, string(char), "", -1)
	}

	return correctedBodyString
}

func toFeedItem(i *gofeed.Item) *feedItem {
	o := new(feedItem)

	o.Title = i.Title
	o.Description = i.Description

	if i.Image != nil {
		o.ImageURL = i.Image.URL
		o.ImageTitle = i.Image.Title
	}

	o.PublishedDateString = i.Published
	o.PublishedDate = *i.PublishedParsed

	if len(i.Categories) > 0 {
		o.Category = i.Categories[0]
	}

	o.Content = i.Content
	o.SiteURL = i.Link
	o.OriginalItem = i
	return o
}
