package databot

// Record is a collection Field s
type Record struct {
	Fields []*Field `toml:"field"`
}

// RecordCollector creates one or more records from given record info
type RecordCollector interface {
	Collect() []*map[string]*interface{}
}
