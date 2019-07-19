package processor

var processorMap = make(map[string]Processor)

type controlMessage string

const (
	endSplit controlMessage = "EndSplit"
)

// Flow contains the data and control channels for processor
type Flow struct {
	Data    <-chan interface{}
	Control chan controlMessage
}

// Processor defines the signature of data processor
type Processor func(input Input, control Control, params map[string]interface{}) Output

func register(identifier string, processor Processor) {
	processorMap[identifier] = processor
}

// Get returns the operators that is mapped to the input identifier
func Get(identifier string) Processor {
	return processorMap[identifier]
}
