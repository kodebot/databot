package databot

// FieldSpec represents a field in a record along with the collector and transformers linked to it
type FieldSpec struct {
	Name             string
	CollectorSpec    *CollectorSpec     `toml:"collector"`
	TransformerSpecs []*TransformerSpec `toml:"transformer"`
}
