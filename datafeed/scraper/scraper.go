package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kodebot/newsfeed/logger"
)

/*Scrape returns scrapped string from html
source - can be url or html string
params - can have
			sourceType - default is url. allowed values are url and html
			selectorType - default is css. allowed values are css, xpath and custom
			selector - selector for chosen selectorType
			all the params for custom selector is prefixed with custom:

*/
func Scrape(source string, params map[string]interface{}) string {

	if source == "" {
		return source
	}

	var sourceType string

	// use url when for source type when type param is missing
	// allowed types are url and html
	if params["sourceType"] == nil {
		sourceType = "url"
	} else {
		sourceType = params["sourceType"].(string)
	}

	// css, xpath and custom are supported
	var selectorType string
	if params["selectorType"] == nil { // default to css
		selectorType = "css"
	} else {
		selectorType = params["selectorType"].(string)
	}

	switch selectorType {
	case "custom":
		var selector string
		if params["custom:selector"] != nil {
			selector = params["custom:selector"].(string)
		}

		var focusSelectors []string
		if params["custom:focusSelectors"] != nil {
			for _, focusSelector := range params["custom:focusSelectors"].([]interface{}) {
				focusSelectors = append(focusSelectors, focusSelector.(string))
			}
		}

		var blacklistedSelectors []string
		if params["custom:blacklistedSelectors"] != nil {
			for _, blacklistedSelector := range params["custom:blacklistedSelectors"].([]interface{}) {
				blacklistedSelectors = append(blacklistedSelectors, blacklistedSelector.(string))
			}
		}

		var fallbackImageSelector string
		if params["custom:fallbackImageSelector"] != nil {
			fallbackImageSelector = params["custom:fallbackImageSelector"].(string)
		}

		switch selector {
		case "content":
			return extractContent(source, sourceType, focusSelectors, blacklistedSelectors, fallbackImageSelector)
		case "guided":
			return extractGuidedContent(source, sourceType, focusSelectors, blacklistedSelectors, fallbackImageSelector)
		default:
			logger.Errorf("the custom selector type %s is not implemented", selector)
		}

	default:
		logger.Errorf("the selector type %s is not implemented", selectorType)
	}

	return ""
}

func removeNodes(s *goquery.Selection) {
	s.Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		if parent.Length() > 0 {
			parent.Get(0).RemoveChild(s.Get(0))
		}
	})
}

func removeEmptyNodes(s *goquery.Selection) {
	//s.Find("p,div,span,ul,li,section").Each(func(i int, s *goquery.Selection) {
	s.Find("*").Not("img,br").Each(func(i int, s *goquery.Selection) {
		if len(s.Find("img,br").Nodes) != 0 {
			return
		}
		if len(strings.TrimSpace(s.Text())) == 0 {
			removeNodes(s)
		}
	})
	//})
}

func stripStyles(s *goquery.Selection) {
	s.Find("*").Each(func(i int, s *goquery.Selection) {
		s.RemoveAttr("style")
	})
}

func stripClasses(s *goquery.Selection) {
	s.Find("*").Each(func(i int, s *goquery.Selection) {
		s.RemoveAttr("class")
	})
}

func removeAdvertisementLeftovers(s *goquery.Selection) {
	s.Find("*").Each(func(i int, s *goquery.Selection) {
		if strings.TrimSpace(s.Text()) == "Advertisement" {
			removeNodes(s)
		}
	})
}

func removeCommentNodes(s *goquery.Selection) {
	s.Contents().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "#comment" {
			removeNodes(s)
		}
	})
}

func removeAllDataDashAttrs(s *goquery.Selection) {
	s.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, a := range s.Get(0).Attr {
			if strings.HasPrefix(a.Key, "data-") {
				s.RemoveAttr(a.Key)
			}
		}
	})
}
