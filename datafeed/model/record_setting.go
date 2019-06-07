package model

import (
	collector_model "github.com/kodebot/newsfeed/datafeed/collectors/model"
)

// RecordSetting allows to specified record setting for parsing data
type RecordSetting struct {
	Type          collector_model.RecordCollectorType
	FieldSettings []FieldSetting
}
