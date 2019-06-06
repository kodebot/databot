package model

// FieldCollectorType provides available field collector types
type FieldCollectorType int

const (
	// VALUE field collector - the result will be same as source field
	VALUE FieldCollectorType = iota
	// REGEXP field collector
	REGEXP

	// CSS field collector
	CSS
)
