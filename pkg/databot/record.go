package databot

// RecordSpec is a collection Field s
type RecordSpec struct {
	CollectorSpec *RecordCollectorSpec `toml:"collector"`
	FieldSpecs    []*FieldSpec         `toml:"field"`
}

// RecordCollectorSpec is the spec for record collector
type RecordCollectorSpec struct {
	Type RecordCollectorType
	// Source is usually URL - other sources like file, etc... are not supported
	SourceURI string
	Params    map[string]interface{}
}

// RecordCollectorType provides available collectors
type RecordCollectorType string

const (
	// RssAtomRecordCollector represents the rss/atom record collector
	RssAtomRecordCollector RecordCollectorType = "rssatom"
	// HTMLSingleRecordCollector represents single html record collector
	HTMLSingleRecordCollector = "htmlSingle"
	// HTMLRecordCollector represents multiple html record collector
	HTMLRecordCollector = "html"
)

// todo: introduce record collector and field collector when needed
// idea for record collector: on htmlmultiple source type - collecting the link to collect individual record fields can be record collector
// for example on news website home page, we will harvest the news item links where each news item link is a record (more like a skeleton record).
// The record collector will collect these links and pass them on to the field collector. We can have something like record pre collector to get html content
// from the links
