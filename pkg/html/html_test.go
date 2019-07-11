package html

import (
	"testing"

	"github.com/kodebot/databot/pkg/databot"
)

func TestCollectRecord(t *testing.T) {

	// reader := NewDocumentReader("https://www.dailythanthi.com/News/India")
	reader := NewDocumentReader("https://www.bbc.co.uk/news")
	src, _ := reader.ReadAsString()

	// #nw-c-topstories-domestic
	// .gs-c-promo-body a.gs-c-promo-heading

	spec := databot.RecordCollectorSpec{
		Params: map[string]interface{}{
			// "css:remove":        []string{"head"},
			// "css:select":        []string{"body"},
			// "css:selectEach":    []string{".ListingNewsWithMEDImage a:first-child"},
			// "regexp:remove":     []string{"(?P<data>^<a)"},
			// "regexp:select":     []string{"href=\"(?P<data>[^\"]*)"},
			// "regexp:selectEach": []string{"(?P<data>[^/]*)"},

			"css:select":     []string{"#nw-c-topstories-domestic"},
			"css:selectEach": []string{".gs-c-promo-body a.gs-c-promo-heading"},
			"regexp:select":  []string{"<a.*?>(?P<data>.*?)</a>"},
			"fetch":          true,
		},
	}
	result := collectRecord(src, &spec)
	t.Fatalf("%#v", result)
}
