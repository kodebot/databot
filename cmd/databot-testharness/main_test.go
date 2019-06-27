package testharness

import (
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
)

func Test(t *testing.T) {

	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.Read("feedconfig.toml")

	switch feed.SourceType {
	case databot.RssAtomFeedSource:
		xml, err := html.ReadAsString(feed.SourceURI)
		if err != nil {
		}
		rssAtomFeed := rssatom.Parse(xml)

		engine := rssatom.NewRecordEngine(feed.RecordSpec, rssAtomFeed)
		recs := engine.CreateRecords()

		outputPath := "./result.txt"
		resultStr := toString(recs)
		err = ioutil.WriteFile(outputPath, []byte(resultStr), os.ModePerm)
		if err != nil {
			logger.Fatalf("unable to write to file %s. error: %s", outputPath, err.Error())
		}

	case databot.HTMLSingleFeedSource:
		panic(errors.New("HTMLSingle source is not implemented"))

	case databot.HTMLFeedSource:
		panic(errors.New("HTMLMultiple source is not implemented"))
	default:
		panic(errors.New("Unsupported source"))
	}

}

func toString(records []*map[string]*interface{}) string {
	result := []string{}
	for _, record := range records {
		fields := []string{}
		for key, value := range *record {
			valueString := "NIL"
			if value != nil {
				if valueDate, ok := (*value).(*time.Time); ok {
					valueString = valueDate.String()
				} else if valueDate, ok := (*value).(time.Time); ok {
					valueString = valueDate.String()
				} else {
					valueString = (*value).(string)
				}
			}
			fields = append(fields, key+": "+valueString)
		}
		sort.Strings(fields)
		result = append(result, strings.Join(fields, ",\n"))
	}

	return strings.Join(result, "\n**********************************************************************************\n")
}
