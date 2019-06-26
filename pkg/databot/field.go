package databot

// Field represents a field in a record along with the collector and transformers linked to it
type Field struct {
	Name         string
	Collector    *Collector
	Transformers []*Transformer
}

// FieldCollector collects a field data from given field info
type FieldCollector interface {
	Collect() []*map[string]*interface{}
}
