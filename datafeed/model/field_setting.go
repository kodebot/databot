package model

import (
	cmodel "github.com/kodebot/newsfeed/datafeed/collectors/model"
	tmodel "github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// FieldSetting allows to specify field setting when parsing
type FieldSetting struct {
	Field               string
	CollectorSetting    cmodel.FieldCollectorSetting
	TransformerSettings []tmodel.TransformerSetting
}
