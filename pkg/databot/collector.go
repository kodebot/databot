package databot

// CollectorSpec represents collector config
type CollectorSpec struct {
	Type   CollectorType
	Params *map[string]*interface{}
}

// CollectorType provides available collector types
type CollectorType int

const (
	// Pluck represents a type of collector that fetches value from property of an instance
	Pluck CollectorType = iota + 1
)
