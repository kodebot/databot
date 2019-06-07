package transformers

import (
	"time"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// Transform returns transformed source as per given transformer settings
func Transform(source *interface{}, transformerSettings []model.TransformerSetting) *interface{} {

	for _, transformerSetting := range transformerSettings {
		sourceValue := *source
		switch transformerSetting.Transformer {
		case model.Trim:
			if sourceString, ok := sourceValue.(string); ok {
				sourceValue = trim(sourceString, transformerSetting.Parameters)
			} else {
				glog.Errorf("trim is not allowed on non string type")
			}
		case model.FormatDate:
			if sourceTime, ok := sourceValue.(time.Time); ok {
				sourceValue = formatDate(sourceTime, transformerSetting.Parameters)
			} else {
				glog.Errorf("formatDate is not allowed on non time.Time type")
			}

		default:
			glog.Errorf("the specified transformer %s in not supported", transformerSetting.Parameters)
		}
		source = &sourceValue
	}

	return source
}
