package services

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kodebot/newsfeed/models"
)

func TestParseFeed(t *testing.T) {
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile("../feed_config.toml", &feedConfig)
	if err != nil {
		t.Errorf("error when loading feed config: %s\n", err.Error())
	}

	for _, feed := range feedConfig.Feed {
		t.Logf("processing %s \n", feed.URL)
		result := ParseFeed(feed)
		if result == nil {
			t.Error("failed")
		}
		t.Logf("finished processing %s \n", feed.URL)
	}
}

func xTestIllegalXML(t *testing.T) {

	illegalXMLCharacters := []rune{
		'\u0001', '\u0002', '\u0003', '\u0004', '\u0005', '\u0006', '\u0007',
		'\u0008', '\u000b', '\u000c', '\u000e', '\u000f', '\u0010', '\u0011',
		'\u0012', '\u0013', '\u0014', '\u0015', '\u0016', '\u0017', '\u0018',
		'\u0019', '\u001a', '\u001b', '\u001c', '\u001d', '\u001e', '\u001f'}

	var client http.Client
	resp, err := client.Get("https://sports.dinamalar.com/rss/Badminton")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	t.Errorf("%d", http.StatusOK)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		bodyString := string(bodyBytes)
		t.Log(bodyString)
		correctedBodyString := bodyString

		for _, char := range illegalXMLCharacters {
			correctedBodyString = strings.Replace(correctedBodyString, string(char), "", -1)
		}

		t.Log(correctedBodyString)
	}
}
