package pipeline

var processorMap = make(map[string]Builder)

func register(identifier string, builder Builder) {
	processorMap[identifier] = builder
}

type Builder func(params map[string]interface{}) Operator

type Operator func(in <-chan interface{}, out chan<- interface{})

func Get(identifier string) Builder {
	return processorMap[identifier]
}
