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

func (ctx *htmlContext) removeElements(val interface{}, params map[string]interface{}) interface{} {
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

func (ctx *htmlContext) selectElements(val interface{}, params map[string]interface{}) interface{} {
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

func (ctx *htmlContext) removeStyles(val interface{}, params map[string]interface{}) interface{} {
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

func (ctx *htmlContext) removeScripts(val interface{}, params map[string]interface{}) interface{} {
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
	doc.RemoveAttrsWhen(func(attr *gohtml.Attribute) bool {
		return strings.Contains(attr.Key, "data-") || strings.Contains(attr.Val, "javascript:")
	})
	return doc.HTML()
}

func (ctx *htmlContext) removeNonContentElements(val interface{}, params map[string]interface{}) interface{} {
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

func (ctx *htmlContext) removeElementsMatchingText(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	if params == nil || params["matchers"] == nil {
		logger.Errorf("missing params or no matchers in the params ")
		return nil
	}

	matchers := toStringSlice(params["matchers"])

	if len(matchers) == 0 {
		logger.Errorf("empty matchers found in the params")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	doc.RemoveNodesWhen(func(n *gohtml.Node) bool {
		for _, matcher := range matchers {
			match, err := regexp.MatchString(matcher, n.Data)
			if err == nil && match {
				return true
			}
		}
		return false
	})

	return doc.HTML()
}

func (ctx *htmlContext) getMetadata(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	htmlStr, ok := val.(string)
	if !ok {
		logger.Errorf("input is not a string")
		return nil
	}

	if params == nil {
		logger.Errorf("params must be present with keyAttr, keyVal and valAttr")
		return nil
	}

	keyAttr := params["keyAttr"]
	keyVal := params["keyVal"]
	valAttr := params["valAttr"]

	if keyAttr == nil || keyVal == nil || valAttr == nil {
		logger.Errorf("keyAttr, keyVal and valAttr must be specified for htmlMetadata transformer")
		return nil
	}

	doc := ctx.newDocFn(htmlStr)
	return doc.GetMetadata(keyAttr.(string), keyVal.(string), valAttr.(string))
}

func toStringSlice(slice interface{}) []string {
	if s, ok := slice.([]interface{}); ok {
		var strs []string
		for _, i := range s {
			strs = append(strs, i.(string))
		}
		return strs
	}
	return []string{}
}
