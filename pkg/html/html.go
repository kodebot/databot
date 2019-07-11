package html

import (
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/fldcollector"
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

	recSources := r.getRecordSources(spec.CollectorSpec)
	collected := collect(recSources, spec)
	// todo: review whether it is ok to collect all records and transform or we need to collect and transform one record at a time
	// todo: no record transformers are supported now
	transformed := applyRecTransformers(collected, nil)
	return transformed
}

func collect(sources []string, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}
	for _, item := range sources {
		rec := make(map[string]interface{})
		for _, fldSpec := range spec.FieldSpecs {
			rec[fldSpec.Name] = createField(item, fldSpec)
		}
		recs = append(recs, "")
	}
	return recs
}

func (r *recordCreator) getRecordSources(collectorSpec *databot.RecordCollectorSpec) []string {
	docReader := r.docReaderFn(collectorSpec.SourceURI)
	source, err := docReader.ReadAsString()
	if err != nil {
		logger.Errorf("unable to retrieve content from URI %s", collectorSpec.SourceURI)
	}

	result := []string{}
	result = append(result, source)

	for key, val := range collectorSpec.Params {
		switch key {
		case "fetch":
			result = fetch(result, val, collectorSpec)
		case "css:remove":
			result = cssRemove(result, val)
		case "css:select":
			result = cssSelect(result, val)
		case "css:selectEach":
			result = cssSelectEach(result, val)
		case "regexp:remove":
			result = regexpRemove(result, val)
		case "regexp:select":
			result = regexpSelect(result, val)
		case "regexp:selectEach":
			result = regexpSelectEach(result, val)
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

func cssRemove(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	for i, block := range input {
		doc := NewDocument(block)
		doc.Remove(selectors...)
		input[i] = doc.HTML()
	}
	return input
}

func cssSelect(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	for i, block := range input {
		doc := NewDocument(block)
		doc.Select(selectors...)
		input[i] = doc.HTML()
	}
	return input
}

func cssSelectEach(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	newResult := []string{}
	for _, block := range input {
		doc := NewDocument(block)
		newResult = append(newResult, doc.HTMLEach(selectors...)...)
	}
	return newResult
}

func regexpRemove(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	for i, block := range input {
		for _, selector := range selectors {
			matches := regexpMatchAll(block, selector)
			for _, match := range matches {
				block = strings.Replace(block, match, "", -1)
			}
		}
		input[i] = block
	}
	return input
}

func regexpSelect(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	for i, block := range input {
		for _, selector := range selectors {
			matches := regexpMatchAll(block, selector)
			block = strings.Join(matches, "")
		}
		input[i] = block
	}
	return input
}

func regexpSelectEach(input []string, param interface{}) []string {
	selectors, ok := param.([]string)
	if !ok {
		panic("selector must be specified using slice of string")
	}
	newResult := []string{}
	for _, block := range input {
		for _, selector := range selectors {
			matches := regexpMatchAll(block, selector)
			newResult = append(newResult, matches...)
		}
	}
	return newResult
}

func fetch(input []string, param interface{}, spec *databot.RecordCollectorSpec) []string {
	if param != nil && param.(bool) {
		result := []string{}
		for i, url := range input {
			url = resolveRelativePath(spec.SourceURI, url)
			docReader := NewCachedDocumentReader(url, cache.Current())
			htmlStr, err := docReader.ReadAsString()
			if err != nil {
				logger.Errorf("unable to get html from url: %s", url)
				result = append(result, "")
			} else {
				result = append(result, htmlStr)
			}
		}
		return result
	}
	return input
}

func createField(source string, spec *databot.FieldSpec) interface{} {
	collected := collectField(source, spec.CollectorSpec)
	return transformField(collected, spec.TransformerSpecs)
}

func collectField(source string, spec *databot.FieldCollectorSpec) interface{} {
	collectorType := spec.Type

	// if collector := collectorMap[collectorType]; collector != nil {
	// 	return collector(source, spec.Params)
	// }

	if sharedCollector := fldcollector.CollectorMap[collectorType]; sharedCollector != nil {
		var source interface{} = source
		return sharedCollector(&source, spec.Params)
	}

	logger.Errorf("specified collector %s is not found", collectorType)
	return nil
}

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
