package processor

var processorMap = make(map[string]Builder)

func register(identifier string, builder Builder) {
	processorMap[identifier] = builder
}

// Processor is a data processor function that can be applied to transform input data
type Processor func(in <-chan interface{}, out chan<- interface{})

// Builder defines a function that returns Processor
type Builder func(params map[string]interface{}) Processor

// GetProcessorBuilder returns Builder that matches the identifier
func GetProcessorBuilder(identifier string) Builder {
	return processorMap[identifier]
}
