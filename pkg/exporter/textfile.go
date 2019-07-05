package exporter

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/kodebot/databot/pkg/logger"
)

// ExportToTextFile exports the records into text file in a readable format
func ExportToTextFile(records []map[string]interface{}, filePath string) {
	result := []string{}
	for _, record := range records {
		fields := []string{}
		for key, value := range record {
			valueString := "NIL"
			switch v := value.(type) {
			case *time.Time:
				valueString = v.String()
			case time.Time:
				valueString = v.String()
			case string:
				valueString = v
			case *interface{}:
				valueString = fmt.Sprintf("%+v", *v)
			default:
				valueString = fmt.Sprintf("%+v", v)
			}

			fields = append(fields, key+": "+valueString)
		}
		sort.Strings(fields)
		result = append(result, strings.Join(fields, ",\n"))
	}

	resultStr := strings.Join(result, "\n**********************************************************************************\n")

	err := ioutil.WriteFile(filePath, []byte(resultStr), os.ModePerm)
	if err != nil {
		logger.Fatalf("unable to write to file %s. error: %s", filePath, err.Error())
	}
}
