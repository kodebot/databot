package databot

// RecordSpec is a collection Field s
type RecordSpec struct {
	FieldSpecs []*FieldSpec `toml:"field"`
}
