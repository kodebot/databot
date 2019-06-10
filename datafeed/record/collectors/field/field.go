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
	result = fcollectors.Collect(source, info.CollectorInfo)
	return ftransformers.Transform(result, info.TransformersInfo)
}
