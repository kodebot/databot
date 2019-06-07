package model

import (
	cmodel "github.com/kodebot/newsfeed/datafeed/collectors/model"
)

// RecordSetting allows to specified record setting for parsing data
type RecordSetting struct {
	Type          cmodel.RecordCollectorType
	Source        string
	FieldSettings []FieldSetting `toml:"field"`
}
