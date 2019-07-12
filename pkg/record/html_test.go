package record

import (
	"testing"

	"github.com/kodebot/databot/pkg/config"

	"github.com/kodebot/databot/pkg/fldtransformer"

	"github.com/kodebot/databot/pkg/databot"
)

func TestCollectRecord(t *testing.T) {
	config.NewBuilder().Build()
	spec := databot.RecordSpec{
		CollectorSpec: &databot.RecordCollectorSpec{
			SourceURI: "https://www.bbc.co.uk/news",
			Params: map[string]interface{}{
				// "css:remove":        []string{"head"},
				// "css:select":        []string{"body"},
				// "css:selectEach":    []string{".ListingNewsWithMEDImage a:first-child"},
				// "regexp:remove":     []string{"(?P<data>^<a)"},
				// "regexp:select":     []string{"href=\"(?P<data>[^\"]*)"},
				// "regexp:selectEach": []string{"(?P<data>[^/]*)"},

				"css:select":     []string{"#nw-c-topstories-domestic"},
				"css:selectEach": []string{".gs-c-promo-body a.gs-c-promo-heading"},
				"regexp:select":  []string{"href=\"(?P<data>[^\"]*)"},
				"fetch":          true,
			},
		},
		FieldSpecs: []*databot.FieldSpec{
			&databot.FieldSpec{
				Name: "Title",
				CollectorSpec: &databot.FieldCollectorSpec{
					Type: "source",
				},
				TransformerSpecs: []*databot.FieldTransformerSpec{
					&databot.FieldTransformerSpec{
						Type: fldtransformer.SelectHTMLElements,
						Params: map[string]interface{}{
							"selectors": []string{"h3"},
						},
					},
				}},
		},
	}
	result := NewRecordCreator().Create(&spec)
	t.Fatalf("%#v", result)
}
