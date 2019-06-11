package record

import (
	"testing"

	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
)

var recordCollectMockCalled = false

func TestCreateCallRssAtomCollect(t *testing.T) {
	rcollect = recordCollectMock
	result := Create("", rcollectors.RssAtom, *new(Info))
	if recordCollectMockCalled != true {
		t.Errorf("expect record collector to be called but it is not called")
	}

	if len(result) != 1 {
		t.Errorf("expected result is not returned from RssAtom collector")
	}

	resetMockData()
}

func recordCollectMock(data string, sourceType rcollectors.SourceType, fields []field.Info) []map[string]interface{} {
	recordCollectMockCalled = true
	return []map[string]interface{}{{"foo": "bar"}}
}

func resetMockData() {
	recordCollectMockCalled = false
}
