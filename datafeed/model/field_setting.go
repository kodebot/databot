package model

import (
	"github.com/kodebot/newsfeed/datafeed/collectors/record/fields"
	tmodel "github.com/kodebot/newsfeed/datafeed/transformers/model"
)

// FieldInfo allows to specify field setting when parsing
type FieldInfo struct {
	Name                string
	CollectorSetting    fields.CollectorInfo
	TransformerSettings []tmodel.TransformerSetting
}
