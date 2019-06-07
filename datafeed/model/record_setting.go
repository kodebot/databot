package model

// RecordSetting allows to specified record setting for parsing data
type RecordSetting struct {
	FieldSettings []FieldSetting `toml:"field"`
}
