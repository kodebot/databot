package record

import (
	"testing"

	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
)

var rssAtomCollectCalled = false

func TestCreateCallRssAtomCollect(t *testing.T) {
	rssAtomCollect = rssAtomCollectMock
	result := Create("", rcollectors.RssAtom, *new(Info))
	if rssAtomCollectCalled != true {
		t.Errorf("expect RssAtom collector to be called but it is not called")
	}

	if len(result) != 1 {
		t.Errorf("expected result is not returned from RssAtom collector")
	}

	resetMockData()
}

func TestCreateFailsWhenCollectorIsUnknown(t *testing.T) {
	rssAtomCollect = rssAtomCollectMock
	result := Create("", "", *new(Info))
	if rssAtomCollectCalled == true {
		t.Errorf("expect RssAtom collector NOT to be called but it is called")
	}

	if len(result) != 0 {
		t.Errorf("expected result is not returned from RssAtom collector")
	}

	resetMockData()
}

func rssAtomCollectMock(data string, fields []field.Info) []map[string]interface{} {
	rssAtomCollectCalled = true
	return []map[string]interface{}{{"foo": "bar"}}
}

func resetMockData() {
	rssAtomCollectCalled = false
}
