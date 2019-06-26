package databot

// Collector represents collector config
type Collector struct {
	Type   *CollectorType
	Params *map[string]*interface{}
}

// CollectorType provides available collector types
type CollectorType string
