package field

import (
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	ghtml "golang.org/x/net/html"
)

func removeHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	if htmlStr, ok := val.(string); ok {
		if found := params["selectors"]; found != nil {
			selectors := toStringSlice(found)
			if len(selectors) > 0 {
				doc := html.NewDocument(htmlStr)
				doc.Remove(selectors...)
				return doc.HTML()
			}
		}
		logger.Errorf("no selectors found")
		return val
	}

	logger.Errorf("input is not a string")
	return val
}

func selectHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	if htmlStr, ok := val.(string); ok {
		if found := params["selectors"]; found != nil {
			selectors := toStringSlice(found)
			if len(selectors) > 0 {
				doc := html.NewDocument(htmlStr)
				doc.Select(selectors...)
				return doc.HTML()
			}
		}
		logger.Errorf("no selectors found")
		return val
	}

	logger.Errorf("input is not a string")
	return val
}

func removeHTMLStyles(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return val
	}

	doc := html.NewDocument(htmlStr)
	doc.Remove("style")
	doc.RemoveAttrs("style", "class")
	return doc.HTML()
}

func removeHTMLScripts(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return val
	}

	doc := html.NewDocument(htmlStr)
	doc.Remove("script")
	doc.RemoveAttrsWhen(func(attr string, val string) bool {
		return strings.Contains(attr, "data-") || strings.Contains(val, "javascript:")
	})
	return doc.HTML()
}

func removeNonContentHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return val
	}

	doc := html.NewDocument(htmlStr)
	doc.RemoveNonContent()
	return doc.HTML()
}

func removeHTMLElementsMatchingText(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return val
	}

	doc := html.NewDocument(htmlStr)
	doc.RemoveNodeWhen(func(n *ghtml.Node) bool {
		if found := params["matchers"]; found != nil {
			matchers := toStringSlice(found)
			for _, matcher := range matchers {
				match, err := regexp.MatchString(matcher, n.Data)
				if err == nil && match {
					return true
				}
			}
		}
		logger.Errorf("no matchers/matched nodes found")
		return false
	})

	return doc.HTML()
}

func toStringSlice(selectors interface{}) []string {
	if s, ok := selectors.([]interface{}); ok {
		var selectorStrs []string
		for _, selector := range s {
			selectorStrs = append(selectorStrs, selector.(string))
		}
		return selectorStrs
	}
	return []string{}
}
