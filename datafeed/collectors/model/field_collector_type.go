package model

// FieldCollectorType provides available field collector types
type FieldCollectorType string

const (
	// Value field collector - the result will be same as source field
	Value FieldCollectorType = "value"
	// Regexp field collector
	Regexp FieldCollectorType = "regexp"

	// CSS field collector
	CSS FieldCollectorType = "css"

	// Unknown field collector
	Unknown FieldCollectorType = "unknown"
)
