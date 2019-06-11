package collectors

import "testing"

var rssAtomCollectorMockCalled = false

func TestCollectWithValidCollector(t *testing.T) {
	collectorsMap[RssAtomField] = rssAtomCollectorMock
	result := Collect("", CollectorInfo{Type: RssAtomField})

	if rssAtomCollectorMockCalled != true {
		t.Errorf("expect collector to be called but not")
	}

	if result != "foo" {
		t.Errorf("collector doesn't return correct value")
	}
	resetMock()
}

func TestCollectWithInvalidCollector(t *testing.T) {
	result := Collect("", CollectorInfo{Type: "Unknown"})

	if rssAtomCollectorMockCalled == true {
		t.Errorf("expect no collector to be called but it does")
	}

	if result != nil {
		t.Errorf("collector should return nil but doesn't")
	}
	resetMock()
}

func rssAtomCollectorMock(value interface{}, parameters map[string]interface{}) interface{} {
	rssAtomCollectorMockCalled = true
	return "foo"
}

func resetMock() {
	rssAtomCollectorMockCalled = false
}
