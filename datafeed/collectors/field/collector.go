package field

// collector behaviors
type collector interface {
	Collect() string
}

// Collector is generic field collector
type Collector struct {
	collector
	Field      interface{}
	Parameters map[string]interface{}
}
