package rssatom

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/kodebot/databot/pkg/fldcollector"
	"github.com/kodebot/databot/pkg/reccollector"
	"github.com/stretchr/testify/assert"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
)

type testDocumentReader struct {
	returnData  string
	returnError bool
}

func (r *testDocumentReader) ReadAsString() (string, error) {
	if r.returnError {
		return "", errors.New("foo")
	}
	return r.returnData, nil
}

func newTestRecordCreator(t *testing.T, filePath string, returnError bool) databot.RecordCreator {
	resp, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return &recordCreator{
		docReaderFn: func(url string, cacheManger cache.Manager) html.DocumentReader {
			return &testDocumentReader{string(resp), returnError}
		},
		cacheManagerFn: cache.NewMemCache}
}

func TestCreate(t *testing.T) {

	tests := map[string]struct {
		filePath string
		spec     *databot.RecordSpec
		expected []map[string]interface{}
	}{
		"simple rss smoke": {
			filePath: "testdata/rss/rss_foreign_lang.xml",
			spec: &databot.RecordSpec{
				CollectorSpec: &databot.RecordCollectorSpec{SourceURI: "foo", Type: reccollector.RssAtom},
				FieldSpecs: []*databot.FieldSpec{
					{
						Name:          "Title",
						CollectorSpec: &databot.FieldCollectorSpec{Type: fldcollector.PluckField},
					},
				},
			},
			expected: []map[string]interface{}{
				{"Title": "மத்திய பட்ஜெட்டில் சாமானியருக்கான சலுகைகள் ?"},
				{"Title": "'மேற்கு வங்கத்தின் பெயரை 'பங்க்ளா' என மாற்ற முடியாது'"},
			},
		},
	}

	for name, test := range tests {
		creator := newTestRecordCreator(t, test.filePath, false)
		actual := creator.Create(test.spec)
		assert.Equal(t, test.expected, actual, name)
	}
}
