package model

// FieldCollectorSetting settings for collecting field
type FieldCollectorSetting struct {
	Field      string
	Type       FieldCollectorType
	Parameters map[string]interface{}
}
