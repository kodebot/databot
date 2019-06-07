package datafeed

import (
	"github.com/kodebot/newsfeed/datafeed/collectors"
	collectors_model "github.com/kodebot/newsfeed/datafeed/collectors/model"
	"github.com/kodebot/newsfeed/datafeed/model"
	"github.com/kodebot/newsfeed/datafeed/transformers"
	transformers_model "github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// Parse the data from the given URL and provides them as structured map
func Parse(url string, setting model.RecordSetting) []map[string]interface{} {

	var fieldCollectorSettings []collectors_model.FieldCollectorSetting
	var fieldTransformerSettingsMap map[string][]transformers_model.TransformerSetting

	for _, fieldSetting := range setting.FieldSettings {
		fieldSetting.CollectorSetting.Field = fieldSetting.Field
		fieldCollectorSettings = append(fieldCollectorSettings, fieldSetting.CollectorSetting)

		fieldTransformerSettingsMap[fieldSetting.Field] = fieldSetting.TransformerSettings
	}

	collectedRecords := collectors.Collect(url, setting.Type, fieldCollectorSettings)

	for _, record := range collectedRecords {
		for fieldName, fieldVal := range record {
			if fieldTransformerSettings := fieldTransformerSettingsMap[fieldName]; fieldCollectorSettings != nil {
				record[fieldName] = transformers.Transform(fieldVal, fieldTransformerSettings)
			}
		}
	}

	return collectedRecords
}
