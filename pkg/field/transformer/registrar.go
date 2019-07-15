package transformer

var transformerMap = make(map[string]Transformer)

// Transformer defines the signature of all the field transformers
type Transformer func(<-chan interface{}, map[string]interface{}) <-chan interface{}

func register(identifier string, transformer Transformer) {
	transformerMap[identifier] = transformer
}

// Get returns the operators that is mapped to the input identifier
func Get(identifier string) Transformer {
	return transformerMap[identifier]
}
