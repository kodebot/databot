package html

import (
	"strings"

	"github.com/kodebot/databot/pkg/logger"

	"github.com/PuerkitoBio/goquery"
)

// Document represents HTML document model
type Document struct {
	document *goquery.Document
}

// NewDocument returns a model that represents HTML document
func NewDocument(htmlStr string) *Document {
	goqueryDoc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		logger.Errorf("unable to create goquery document. error: %s", err.Error())
	}
	return &Document{goqueryDoc}
}

// Remove removes the elements from the document that matches given selectors
func (d *Document) Remove(selectors ...string) {
	selectorsStr := strings.Join(selectors, ",")
	d.document.Find(selectorsStr).Remove()
}

// Keep only keeps the matching elements in the document
func (d *Document) Keep(selectors ...string) {
	selectorsStr := strings.Join(selectors, ",")
	htmlToKeep := new(strings.Builder)
	d.document.Find(selectorsStr).Each(func(i int, s *goquery.Selection) {
		htmlStr, _ := goquery.OuterHtml(s)
		htmlToKeep.WriteString(htmlStr)
	})

	d.document.Find("*").Remove()
	d.document.SetHtml(htmlToKeep.String())
}

// Body returns the document body as string
func (d *Document) Body() string {
	htmlStr, err := d.document.Find("body").Html()
	if err != nil {
		logger.Errorf("error when getting the body of html document. error: %s", err.Error())
		return ""
	}
	return htmlStr
}
