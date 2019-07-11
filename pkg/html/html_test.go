package html

import (
	"testing"

	"github.com/kodebot/databot/pkg/databot"
)

func TestCollectRecord(t *testing.T) {

	reader := NewDocumentReader("https://www.dailythanthi.com/News/India")
	src, _ := reader.ReadAsString()

	spec := databot.RecordCollectorSpec{
		Params: map[string]interface{}{
			"css:remove":        []string{"head"},
			"css:select":        []string{"body"},
			"css:selectEach":    []string{".ListingNewsWithMEDImage a:first-child"},
			"regexp:remove":     []string{"(?P<data>^<a)"},
			"regexp:select":     []string{"href=\"(?P<data>[^\"]*)"},
			"regexp:selectEach": []string{"(?P<data>[^/]*)"},
		},
	}
	result := collectRecord(src, &spec)
	t.Fatalf("%#v", result)
}
