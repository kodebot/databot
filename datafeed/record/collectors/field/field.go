package field

import (
	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
	ftransformers "github.com/kodebot/newsfeed/datafeed/record/collectors/field/transformers"
)

// Info allows to specify field setting when parsing
type Info struct {
	Name             string
	CollectorInfo    fcollectors.CollectorInfo
	TransformersInfo []ftransformers.TransformerInfo
}

// Create returns collected and transformed value
func Create(source interface{}, info Info) interface{} {
	var result interface{}

	if info.CollectorInfo.Parameters == nil {
		info.CollectorInfo.Parameters = map[string]interface{}{}
	}

	if info.CollectorInfo.Parameters["source"] == nil {
		info.CollectorInfo.Parameters["source"] = info.Name
	}

	result = fcollectors.Collect(source, info.CollectorInfo)
	return ftransformers.Transform(result, info.TransformersInfo)
}
