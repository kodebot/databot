package field

import (
	"github.com/golang/glog"
	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
	ftransformers "github.com/kodebot/newsfeed/datafeed/record/collectors/field/transformers"
)

// Info allows to specify field setting when parsing
type Info struct {
	Name             string
	CollectorInfo    fcollectors.CollectorInfo
	TransformersInfo []ftransformers.TransformerInfo
}

// Create returns collected, transformed value
func Create(source interface{}, info Info) interface{} {

	var result interface{}

	collector := info.CollectorInfo
	switch collector.Type {
	case fcollectors.Value:
		result = fcollectors.CollectValue(source, collector.Parameters)
	case fcollectors.Regexp:
		result = fcollectors.CollectRegexp(source, collector.Parameters)
	default:
		glog.Errorf("collector type %s is not implemented", collector.Type)
		result = nil
		break
	}

	for _, tInfo := range info.TransformersInfo {
		switch tInfo.Transformer {
		case ftransformers.Trim:
			result = ftransformers.TransformTrim(result, tInfo.Parameters)
		case ftransformers.FormatDate:
			result = ftransformers.TransformFormatDate(result, tInfo.Parameters)
		default:
			glog.Errorf("the specified transformer %s in not supported", tInfo.Transformer)
		}
	}
	return result
}
