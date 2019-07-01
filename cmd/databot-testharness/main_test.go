package testharness

import (
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	recordcollector "github.com/kodebot/databot/pkg/collectors/record"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
)

func Test(t *testing.T) {

	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.Read("feedconfig.toml")

	switch feed.RecordSpec.CollectorSpec.Type {
	case recordcollector.RssAtom:
		factory := rssatom.NewRecordFactory(feed.RecordSpec)
		recs := factory.Create()
		outputPath := "./result.txt"
		resultStr := toString(recs)
		err := ioutil.WriteFile(outputPath, []byte(resultStr), os.ModePerm)
		if err != nil {
			logger.Fatalf("unable to write to file %s. error: %s", outputPath, err.Error())
		}

	case recordcollector.HTMLSingle:
		panic(errors.New("HTMLSingle record collector is not implemented"))

	case recordcollector.HTML:
		panic(errors.New("HTMLMultiple record collector is not implemented"))
	default:
		panic(errors.New("Unsupported record collector"))
	}
}

func toString(records []map[string]interface{}) string {
	result := []string{}
	for _, record := range records {
		fields := []string{}
		for key, value := range record {
			valueString := "NIL"
			if value != nil {
				if valueDate, ok := value.(*time.Time); ok {
					valueString = valueDate.String()
				} else if valueDate, ok := value.(time.Time); ok {
					valueString = valueDate.String()
				} else {
					valueString = value.(string)
				}
			}
			fields = append(fields, key+": "+valueString)
		}
		sort.Strings(fields)
		result = append(result, strings.Join(fields, ",\n"))
	}

	return strings.Join(result, "\n**********************************************************************************\n")
}
