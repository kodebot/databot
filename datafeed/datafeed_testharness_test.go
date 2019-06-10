// build testharness

package datafeed

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHarnessDataFeedFromFeedInfo(t *testing.T) {

	feedConfigPath := "./test_data/testharness/ready/"
	files, _ := ioutil.ReadDir(feedConfigPath)

	for _, file := range files {
		fullPath := feedConfigPath + file.Name()
		dataFeed, _ := NewFromFeedInfo(fullPath)

		outputPath := fullPath + ".txt"
		dataFeedString := dataFeedToString(dataFeed)
		ioutil.WriteFile(outputPath, []byte(dataFeedString), os.ModePerm)

		// todo: delete old file
		// sort the result by key
	}
}

func dataFeedToString(feed []map[string]interface{}) string {
	result := ""
	for _, record := range feed {
		recordString := ""
		for key, value := range record {
			valueString := "NIL"
			if value != nil {
				valueString = value.(string)
			}
			recordString = recordString + "\n" + key + ":" + valueString
		}
		result = result + "\n***********************************************************************************************************\n"
		result = result + "***********************************************************************************************************\n\n"
		result = result + recordString
	}

	return result
}
