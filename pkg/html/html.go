package html

import (
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
)

type recordCreator struct {
	docReaderFn func(string) DocumentReader
}

// NewRecordCreator returns a new record creator that enables creating one or more records using RSS/Atom feed
func NewRecordCreator() databot.RecordCreator {
	return &recordCreator{docReaderFn: NewDocumentReader}
}

// Create returns one or more records from given rss/atom record spec
func (r *recordCreator) Create(spec *databot.RecordSpec) []map[string]interface{} {

	recs := r.collectRecord(spec.CollectorSpec)
	// todo: review whether it is ok to collect all records and transform or we need to collect and transform one record at a time
	// todo: no record transformers are supported now
	// transformed := applyRecTransformers(collected, nil)
	return recs
}

func collect(source string, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}
	//recUrls := collectRecord(source, spec.CollectorSpec)
	// for _, item := range recUrls { // nothing to collect at record level - the feed item is already available
	// rec := make(map[string]interface{})
	// for _, fldSpec := range spec.FieldSpecs {
	// 	normaliseFieldSpec(fldSpec)
	// 	rec[fldSpec.Name] = createField(item, fldSpec)
	// }
	// recs = append(recs, "")
	// }
	return recs
}

func (r *recordCreator) collectRecord(collectorSpec *databot.RecordCollectorSpec) []string {
	docReader := r.docReaderFn(collectorSpec.SourceURI)
	source, err := docReader.ReadAsString()
	if err != nil {
		logger.Errorf("unable to retrieve content from URI %s", collectorSpec.SourceURI)
	}

	result := []string{}
	result = append(result, source)

	for key, val := range collectorSpec.Params {
		if key == "fetch" && val != nil && val.(bool) {
			for i, url := range result {
				url = resolveRelativePath(collectorSpec.SourceURI, url)
				docReader := NewCachedDocumentReader(url, cache.Current())
				htmlStr, err := docReader.ReadAsString()
				if err != nil {
					logger.Errorf("unable to get html from url: %s", url)
					result[i] = ""
					continue
				}
				result[i] = htmlStr
			}
			continue
		}

		keyParts := strings.Split(key, ":")

		if len(keyParts) != 2 {
			panic("invalid key name") // todo: should panic or be more gentle?
		}

		selectorType := keyParts[0]
		action := keyParts[1]

		if selectorType != "css" && selectorType != "regexp" {
			panic("unsupported selector type") // todo: should panic or be more gentle?
		}

		if action != "remove" && action != "select" && action != "selectEach" {
			panic("unsupported action type")
		}

		selectors, ok := val.([]string)
		if !ok {
			panic("selector must be specified using slice of string")
		}

		switch key {
		case "css:remove":
			for i, block := range result {
				doc := NewDocument(block)
				doc.Remove(selectors...)
				result[i] = doc.HTML()
			}

		case "css:select":
			for i, block := range result {
				doc := NewDocument(block)
				doc.Select(selectors...)
				result[i] = doc.HTML()
			}

		case "css:selectEach":
			newResult := []string{}
			for _, block := range result {
				doc := NewDocument(block)
				newResult = append(newResult, doc.HTMLEach(selectors...)...)
			}
			result = newResult

		case "regexp:remove":
			for i, block := range result {
				for _, selector := range selectors {
					matches := regexpMatchAll(block, selector)
					for _, match := range matches {
						block = strings.Replace(block, match, "", -1)
					}
				}
				result[i] = block
			}
		case "regexp:select":
			for i, block := range result {
				for _, selector := range selectors {
					matches := regexpMatchAll(block, selector)
					block = strings.Join(matches, "")
				}
				result[i] = block
			}

		case "regexp:selectEach":
			newResult := []string{}
			for _, block := range result {
				for _, selector := range selectors {
					matches := regexpMatchAll(block, selector)
					newResult = append(newResult, matches...)
				}

			}
			result = newResult
		}
	}
	return result
}

func regexpMatchAll(val string, expr string) []string {
	result := []string{}
	if val != "" {
		if expr == "" {
			logger.Errorf("no regular expression found")
			return result
		}

		re, err := regexp.Compile(expr)
		if err != nil {
			logger.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
			return result
		}

		requiredMatchIndex := -1
		for i, val := range re.SubexpNames() {
			if val == "data" {
				requiredMatchIndex = i
			}
		}

		if requiredMatchIndex > -1 {
			matches := re.FindAllStringSubmatch(val, -1)

			if matches == nil || len(matches) < 1 {
				logger.Warnf("no match found.")
			}

			for _, m := range matches {
				if len(m) < requiredMatchIndex+1 {
					logger.Warnf("no match found.")
					return result
				}

				result = append(result, m[requiredMatchIndex])
			}
		} else { // when there is no captured group 'data' - just return the whole match
			result = re.FindAllString(val, -1)
		}
	}

	return result
}

func resolveRelativePath(sourceURL string, relativePath string) string {
	if strings.HasPrefix(relativePath, "http") {
		return relativePath
	} else if strings.HasPrefix(relativePath, "/") {
		baseURL := regexp.MustCompile("^.+?[^/:](?=[?/]|$)").FindString(sourceURL)
		return baseURL + relativePath
	} else if strings.HasSuffix(sourceURL, "/") {
		return sourceURL + relativePath
	} else {
		return sourceURL + "/" + relativePath
	}
}

func (r *recordCreator) fetchProcessor(input []string, param interface{}) []string {
	if param != nil && param.(bool) {
		for i, url := range input {
			url = resolveRelativePath(r.docReaderFn.SourceURI, url)
			docReader := NewCachedDocumentReader(url, cache.Current())
			htmlStr, err := docReader.ReadAsString()
			if err != nil {
				logger.Errorf("unable to get html from url: %s", url)
				result[i] = ""
				continue
			}
			result[i] = htmlStr
		}
		continue
	}
	return input

}

// func normaliseFieldSpec(field *databot.FieldSpec) {
// 	// initialise params if nil
// 	if params := field.CollectorSpec.Params; params == nil {
// 		field.CollectorSpec.Params = make(map[string]interface{})
// 	}

// 	// set source same as name if missing
// 	if src := field.CollectorSpec.Params["source"]; src == nil {
// 		field.CollectorSpec.Params["source"] = field.Name
// 	}
// }

// func createField(source *gofeed.Item, spec *databot.FieldSpec) interface{} {
// 	if source == nil {
// 		logger.Errorf("Cannot collect field value when RssAtomItem is nil")
// 		return nil
// 	}

// 	collected := collectField(source, spec.CollectorSpec)
// 	return transformField(collected, spec.TransformerSpecs)
// }

// func collectField(source *gofeed.Item, spec *databot.FieldCollectorSpec) interface{} {
// 	collectorType := spec.Type

// 	// for RSS/Atom feed set the collector type to Pluck if not specified
// 	if collectorType == "" {
// 		collectorType = fldcollector.PluckField
// 	}

// 	if collector := collectorMap[collectorType]; collector != nil {
// 		return collector(source, spec.Params)
// 	}

// 	if sharedCollector := fldcollector.CollectorMap[collectorType]; sharedCollector != nil {
// 		var source interface{} = *source
// 		return sharedCollector(&source, spec.Params)
// 	}

// 	logger.Errorf("specified collector %s is not found", collectorType)
// 	return nil
// }

// func transformField(value interface{}, specs []*databot.FieldTransformerSpec) interface{} {
// 	for _, spec := range specs {
// 		if transformerFn := transformersMap[spec.Type]; transformerFn != nil {
// 			value = transformerFn(value, spec.Params)
// 			continue
// 		}

// 		if sharedTransformerFn := fldtransformer.TransformersMap[spec.Type]; sharedTransformerFn != nil {
// 			value = sharedTransformerFn(value, spec.Params)
// 			continue
// 		}

// 		logger.Errorf("transformer %s not found", spec.Type)
// 	}
// 	return value
// }
