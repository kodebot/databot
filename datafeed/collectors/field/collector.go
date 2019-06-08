package field

// collector behaviors
type collector interface {
	Collect() interface{}
}

// Collector is generic field collector
type Collector struct {
	collector
	Field      interface{}
	Parameters map[string]interface{}
}
