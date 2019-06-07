package rssAtom

import (
	"strings"

	"github.com/golang/glog"
	"github.com/mmcdole/gofeed"
)

func readFromXML(xmlString string) []*gofeed.Item {
	glog.Infof("parsing feed xml")
	parser := gofeed.NewParser()
	xmlString = fixIllegalXMLCharacters(xmlString)
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

	return feed.Items
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
