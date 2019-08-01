package processor

var processorMap = make(map[string]Processor)

// ControlMessage that are supported by Flow
type ControlMessage string

const (
	endSplit ControlMessage = "EndSplit"
)

// Flow contains the data and control channels for processor
type Flow struct {
	Data    chan interface{}
	Control chan ControlMessage
}

// Processor defines the signature of data processor
type Processor func(input Flow, params map[string]interface{}) Flow

func register(identifier string, processor Processor) {
	processorMap[identifier] = processor
}

// Get returns the operators that is mapped to the input identifier
func Get(identifier string) Processor {
	return processorMap[identifier]
}
