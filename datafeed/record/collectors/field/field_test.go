package field

import (
	"testing"

	fcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors/field/collectors"
	ftransformers "github.com/kodebot/newsfeed/datafeed/record/collectors/field/transformers"
)

var fcollectMockCollectorInfo fcollectors.CollectorInfo

var fieldCollectorCallOrder = 0
var fieldTransformerCallOrder = 0

func TestCreateCopiesFieldNameAsSourceParamWhenMissing(t *testing.T) {
	tests := []struct {
		input    Info
		expected string
	}{
		{Info{Name: "foo", CollectorInfo: fcollectors.CollectorInfo{Parameters: map[string]interface{}{}}}, "foo"},
		{Info{Name: "foo", CollectorInfo: fcollectors.CollectorInfo{}}, "foo"}}

	for _, test := range tests {
		fcollect = fcollectMock
		Create("", test.input)
		if fcollectMockCollectorInfo.Parameters["source"] != test.expected {
			t.Errorf("source parameter not assigned as it should be")
		}
		resetMock()
	}
}

func TestCreateRunsCollectorThenTransformers(t *testing.T) {
	fcollect = fcollectMock
	ftransform = ftransformMock

	Create("", Info{Name: "foo", CollectorInfo: fcollectors.CollectorInfo{}})
	if fieldCollectorCallOrder != 1 {
		t.Errorf("collector is not called before transformers")
	}

	if fieldTransformerCallOrder != 2 {
		t.Errorf("transformers are not called after collector")
	}

	resetMock()
}

func resetMock() {
	fcollectMockCollectorInfo = *new(fcollectors.CollectorInfo)
	fieldCollectorCallOrder = 0
	fieldTransformerCallOrder = 0
}

func fcollectMock(source interface{}, info fcollectors.CollectorInfo) interface{} {
	fcollectMockCollectorInfo = info
	fieldCollectorCallOrder = fieldTransformerCallOrder + 1
	return nil
}

func ftransformMock(source interface{}, info []ftransformers.TransformerInfo) interface{} {
	fieldTransformerCallOrder = fieldCollectorCallOrder + 1
	return nil
}
