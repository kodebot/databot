package databot

// FieldSpec represents a field in a record along with the collector and transformers linked to it
type FieldSpec struct {
	Name             string
	CollectorSpec    *FieldCollectorSpec     `toml:"collector"`
	TransformerSpecs []*FieldTransformerSpec `toml:"transformer"`
}

// FieldCollectorSpec represents collector config
type FieldCollectorSpec struct {
	Type   FieldCollectorType
	Params map[string]interface{}
}

// FieldCollectorType provides available collectors
type FieldCollectorType string

// FieldTransformerSpec represents transformer config
type FieldTransformerSpec struct {
	Type   FieldTransformerType
	Params map[string]interface{}
}

// FieldTransformerType provides available transformer
type FieldTransformerType string

const (
	// PluckFieldCollector represents a type of collector that fetches value from property of an instance
	PluckFieldCollector FieldCollectorType = "pluck"
)
