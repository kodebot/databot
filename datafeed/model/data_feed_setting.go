package model

// DataFeedSetting provides shape for config data used to create data feed
type DataFeedSetting struct {
	SourceName string
	Source     string
	SourceType DataFeedSourceType
	Category   string
	Schedule   string
	Record     RecordInfo
}
