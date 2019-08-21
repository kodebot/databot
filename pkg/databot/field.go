package databot

// FieldSpec represents a field in a record along with the processors that are applied
type FieldSpec struct {
	Name           string
	ProcessorSpecs []*ProcessorSpec `toml:"processor"`
}
