package model

import (
	collector_model "github.com/kodebot/newsfeed/datafeed/collectors/model"
	transformer_model "github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// FieldSetting allows to specify field setting when parsing
type FieldSetting struct {
	Field               string
	CollectorSetting    collector_model.FieldCollectorSetting
	TransformerSettings []transformer_model.TransformerSetting
}
