package processor

var processorMap = make(map[string]Processor)

// Processor defines the signature of data processor
type Processor func(input <-chan interface{}, params map[string]interface{}) <-chan interface{}

func register(identifier string, processor Processor) {
	processorMap[identifier] = processor
}

// Get returns the operators that is mapped to the input identifier
func Get(identifier string) Processor {
	return processorMap[identifier]
}
