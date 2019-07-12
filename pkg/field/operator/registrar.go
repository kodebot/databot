package operator

var operatorMap = make(map[string]Operator)

// Operator defines the signature of all the field operators
type Operator func(<-chan interface{}, map[string]interface{}) <-chan interface{}

func register(identifier string, fieldOperator Operator) {
	operatorMap[identifier] = fieldOperator
}

// Get returns the operators that is mapped to the input identifier
func Get(identifier string) Operator {
	return operatorMap[identifier]
}
