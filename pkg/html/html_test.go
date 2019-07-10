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
			"html:remover":  []string{""},
			"html:narrower": []string{""},
			"html:selector": []string{".ListingNewsWithMEDImage a:first-child"},
		},
	}
	collectRecord(src, &spec)
}
