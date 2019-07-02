package html

import (
	"strings"

	"github.com/kodebot/databot/pkg/logger"
	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

// Document represents HTML document model
type Document interface {
	Remove(selectors ...string)
	Select(selectors ...string)
	Body() string
	HTML() string
	RemoveAttrs(attrs ...string)
	RemoveAttrsWhen(when func(attr string, val string) bool)
	RemoveNonContent()
	RemoveNodeWhen(when func(node *html.Node) bool)
	GetMetadata(keyAttr string, keyVal string, valAttr string) string
}

type document struct {
	document *goquery.Document
}

// NewDocument returns a model that represents HTML document
func NewDocument(htmlStr string) Document {
	goqueryDoc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		logger.Errorf("unable to create goquery document. error: %s", err.Error())
	}
	doc := document{goqueryDoc}
	return &doc
}

// Remove removes the elements from the document that matches given selectors
func (d *document) Remove(selectors ...string) {
	selectorsStr := strings.Join(selectors, ",")
	d.document.Find(selectorsStr).Remove()
}

// Select only keeps the matching elements in the document
func (d *document) Select(selectors ...string) {
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
func (d *document) Body() string {
	htmlStr, err := d.document.Find("body").Html()
	if err != nil {
		logger.Errorf("error when getting the body of html document. error: %s", err.Error())
		return ""
	}
	return htmlStr
}

// HTML returns the document body as string
func (d *document) HTML() string {
	htmlStr, err := d.document.Html()
	if err != nil {
		logger.Errorf("error when getting the html of the document document. error: %s", err.Error())
		return ""
	}
	return htmlStr
}

// RemoveAttrs removes the specified attribute
func (d *document) RemoveAttrs(attrs ...string) {
	d.document.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, attr := range attrs {
			s.RemoveAttr(attr)
		}
	})
}

// RemoveAttrsWhen removes the specified attribute when the given condition is met
func (d *document) RemoveAttrsWhen(when func(attr string, val string) bool) {
	d.document.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, attr := range s.Get(0).Attr {
			if when(attr.Key, attr.Val) {
				s.RemoveAttr(attr.Key)
			}
		}
	})
}

// RemoveNonContent removes all empty elements including comment elements
func (d *document) RemoveNonContent() {
	d.document.Find("*").Contents().Not("img,br").Each(func(i int, s *goquery.Selection) {
		if len(s.Find("*").Contents().Find("img,br").Nodes) != 0 {
			return
		}
		if len(strings.TrimSpace(s.Text())) == 0 || goquery.NodeName(s) == "#comment" {
			removeNodes(s)
		}
	})
}

// RemoveNodeWhen removes all the nodes matching the given condition
func (d *document) RemoveNodeWhen(when func(node *html.Node) bool) {
	nodesToRemove := []*html.Node{}
	d.document.Find("*").Contents().Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Contents().Nodes {
			if when(node) {
				nodesToRemove = append(nodesToRemove, node)
			}
		}
	})

	s1 := d.document.Find("*").Contents().FilterNodes(nodesToRemove...)
	removeNodes(s1)
}

// GetMetadata retrives the value of valAttr attribute from meta element that that has keyAttr attribute with valAttr value
func (d *document) GetMetadata(keyAttr string, keyVal string, valAttr string) string {
	value := ""
	d.document.Find("meta").Each(func(i int, s *goquery.Selection) {
		actualKeyVal, exist := s.Attr(keyAttr)
		if exist && actualKeyVal == keyVal {
			value = s.AttrOr(valAttr, "")
		}
	})
	return value
}

func removeNodes(s *goquery.Selection) {
	s.Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		if parent.Length() > 0 {
			parent.Get(0).RemoveChild(s.Get(0))
		}
	})
}
