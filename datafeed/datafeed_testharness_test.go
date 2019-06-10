// build testharness

package datafeed

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestHarnessDataFeedFromFeedInfo(t *testing.T) {
	feedConfigPath := "./test_data/testharness/ready/"
	outputDir := "./test_data/testharness/output/"

	err := os.RemoveAll(outputDir)
	if err != nil {
		t.Fatalf("unable to delete the output dir. error %s", err.Error())
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			t.Fatalf("unable to create output dir %s. error: %s", outputDir, err.Error())
		}
	} else {
		t.Fatalf("unable to check if the output dir exist or not. error: %s", err.Error())
	}

	files, err := ioutil.ReadDir(feedConfigPath)
	if err != nil {
		t.Fatalf("unable to read datafeed info from %s. error: %s", feedConfigPath, err.Error())
	}

	for _, file := range files {
		fullPath := feedConfigPath + file.Name()
		dataFeed, _ := NewFromFeedInfo(fullPath)

		outputPath := outputDir + file.Name() + ".txt"
		dataFeedString := dataFeedToString(dataFeed)
		err = ioutil.WriteFile(outputPath, []byte(dataFeedString), os.ModePerm)
		if err != nil {
			t.Fatalf("unable to write to file %s. error: %s", outputPath, err.Error())
		}
	}
}

func dataFeedToString(feed []map[string]interface{}) string {
	result := []string{}
	for _, record := range feed {
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
