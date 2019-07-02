package fldxfmr

import (
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	gohtml "golang.org/x/net/html"
)

type htmlContext struct {
	newDocFn func(string) html.Document
}

func (ctx *htmlContext) removeHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	if htmlStr, ok := val.(string); ok {
		if found := params["selectors"]; found != nil {
			selectors := toStringSlice(found)
			if len(selectors) > 0 {
				doc := ctx.newDocFn(htmlStr)
				doc.Remove(selectors...)
				return doc.HTML()
			}
		}
		logger.Errorf("no selectors found")
		return nil
	}
	logger.Errorf("input is not a string")
	return nil
}

func (ctx *htmlContext) selectHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	if htmlStr, ok := val.(string); ok {
		if found := params["selectors"]; found != nil {
			selectors := toStringSlice(found)
			if len(selectors) > 0 {
				doc := ctx.newDocFn(htmlStr)
				doc.Select(selectors...)
				return doc.HTML()
			}
		}
		logger.Errorf("no selectors found")
		return nil
	}

	logger.Errorf("input is not a string")
	return nil
}

func (ctx *htmlContext) removeHTMLStyles(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	doc.Remove("style")
	doc.RemoveAttrs("style", "class")
	return doc.HTML()
}

func (ctx *htmlContext) removeHTMLScripts(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	doc.Remove("script")
	doc.RemoveAttrsWhen(func(attr string, val string) bool {
		return strings.Contains(attr, "data-") || strings.Contains(val, "javascript:")
	})
	return doc.HTML()
}

func (ctx *htmlContext) removeNonContentHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	doc.RemoveNonContent()
	return doc.HTML()
}

func (ctx *htmlContext) removeHTMLElementsMatchingText(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	doc.RemoveNodeWhen(func(n *gohtml.Node) bool {
		if found := params["matchers"]; found != nil {
			matchers := toStringSlice(found)
			for _, matcher := range matchers {
				match, err := regexp.MatchString(matcher, n.Data)
				if err == nil && match {
					return true
				}
			}
			return false
		}
		logger.Errorf("no matchers/matched nodes found")
		return false
	})

	return doc.HTML()
}

func (ctx *htmlContext) htmlMetadata(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	keyAttr := params["keyAttr"]
	keyVal := params["keyVal"]
	valAttr := params["valAttr"]

	if keyAttr == nil || keyVal == nil || valAttr == nil {
		logger.Errorf("keyAttr, keyVal and valAttr must be specified for htmlMetadata transformer")
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	return doc.GetMetadata(keyAttr.(string), keyVal.(string), valAttr.(string))
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
