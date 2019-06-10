package datafeed

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/collectors"
	"github.com/kodebot/newsfeed/datafeed/collectors/record/fields"
	"github.com/kodebot/newsfeed/datafeed/model"
	"github.com/kodebot/newsfeed/datafeed/transformers"
	tmodel "github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// ParseFromDataFeedSetting returns structured data as per the given data feed settings
func ParseFromDataFeedSetting(filePath string) ([]map[string]interface{}, model.DataFeedSetting) {
	dataFeedSetting := readDataFeedSetting(filePath)
	return ParseFromURL(dataFeedSetting.Source, dataFeedSetting.SourceType, dataFeedSetting.Record), dataFeedSetting
}

// ParseFromURL returns structured data as per the record setting from the given url
func ParseFromURL(url string, sourceType model.DataFeedSourceType, setting model.RecordInfo) []map[string]interface{} {
	data, err := readAsString(url)
	if err != nil {
		glog.Errorf("unable to read from url %s", url)
		return make([]map[string]interface{}, 0)
	}
	return Parse(data, sourceType, setting)
}

// Parse returns structured data as per the record setting from the given data string
func Parse(data string, sourceType model.DataFeedSourceType, setting model.RecordInfo) []map[string]interface{} {
	var fieldCollectorSettings []fields.CollectorInfo
	fieldTransformerSettingsMap := make(map[string][]tmodel.TransformerSetting)

	for _, fieldSetting := range setting.Fields {
		fieldSetting.CollectorSetting.Field = fieldSetting.Name
		fieldCollectorSettings = append(fieldCollectorSettings, fieldSetting.CollectorSetting)

		fieldTransformerSettingsMap[fieldSetting.Name] = fieldSetting.TransformerSettings
	}

	collectedRecords := collectors.Collect(data, sourceType, fieldCollectorSettings)

	for _, record := range collectedRecords {
		for fieldName, fieldVal := range record {
			if fieldTransformerSettings := fieldTransformerSettingsMap[fieldName]; fieldCollectorSettings != nil {
				record[fieldName] = transformers.Transform(fieldVal, fieldTransformerSettings)
			}
		}
	}

	return collectedRecords
}
