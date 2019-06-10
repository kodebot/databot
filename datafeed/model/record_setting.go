package model

// RecordInfo allows to specified record setting for parsing data
type RecordInfo struct {
	Fields []FieldInfo `toml:"field"`
}
