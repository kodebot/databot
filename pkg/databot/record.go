package databot

// RecordSpec is a collection of record processors and field specs
type RecordSpec struct {
	SourceURI         string
	PreprocessorSpecs []*ProcessorSpec `toml:"preprocessor"`
	FieldSpecs        []*FieldSpec     `toml:"field"`
}

// RecordCreator is the abstract record creator
type RecordCreator interface {
	Create(*RecordSpec) []map[string]interface{}
}
