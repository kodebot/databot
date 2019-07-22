package record

import (
	"fmt"
	"sync"

	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/processor"
)

type recordCreator struct {
	docReaderFn func(string) html.DocumentReader
}

// NewRecordCreator returns a new record creator that enables creating one or more records using RSS/Atom feed
func NewRecordCreator() databot.RecordCreator {
	return &recordCreator{docReaderFn: html.NewDocumentReader}
}

// Create returns one or more records from given rss/atom record spec
func (r *recordCreator) Create(spec *databot.RecordSpec) []map[string]interface{} {
	input, output := buildProcessorPipeline(spec.PreprocessorSpecs)

	go func() {
		input.Data <- spec.SourceURI
	}()

	go func() {
		for control := range output.Control {
			// drain
			fmt.Printf("%+v", control)
		}
	}()

	for item := range output.Data {
		fmt.Printf("%+v", item)
	}

	//collected := collect(result.([]interface{}), spec)
	// todo: review whether it is ok to collect all records and transform or we need to collect and transform one record at a time
	// todo: no record transformers are supported now
	// transformed := applyRecTransformers(collected, nil)
	//return collected
	return nil
}

func collect(sources []interface{}, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}

	// fieldProcessorPipelineMap := map[string]processor.Processor

	// for _, fieldSpec := range spec.FieldSpecs {
	//fieldSpec.
	//}

	// for _, item := range sources {
	// 	rec := make(map[string]interface{})
	// 	for _, fldSpec := range spec.FieldSpecs {
	// 		rec[fldSpec.Name] = createField(item, fldSpec)
	// 	}
	// 	recs = append(recs, rec)
	// }
	fmt.Printf("%+v", sources)
	return recs
}

// https://whiskybadger.io/post/introducing-go-pipeline/
func buildProcessorPipeline(processorSpecs []*databot.ProcessorSpec) (processor.Flow, processor.Flow) {
	data := make(chan interface{})
	control := make(chan processor.ControlMessage)

	input := processor.Flow{
		data,
		control,
	}

	output := processor.Flow{}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		pipeline := processor.Flow{}
		pipeline = input
		for _, spec := range processorSpecs {
			print(spec.Name)
			fmt.Printf("%+v", spec.Params)
			nextProcessor := processor.Get(spec.Name)
			pipeline = nextProcessor(pipeline, spec.Params)
		}
		output = pipeline
		wg.Done()

	}()

	wg.Wait()
	return input, output
}

// func buildAndRunProcessorPipeline(initialValue string, processorSpecs []*databot.ProcessorSpec) interface{} {
// 	input := make(chan interface{})
// 	output := []interface{}{}

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		pipeline := make(<-chan interface{})
// 		pipeline = input
// 		for _, spec := range processorSpecs {
// 			print(spec.Name)
// 			fmt.Printf("%+v", spec.Params)
// 			nextProcessor := processor.Get(spec.Name)
// 			pipeline = nextProcessor(pipeline, spec.Params)
// 		}
// 		wg.Done()
// 		for result := range pipeline {
// 			output = append(output, result)
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	wg.Add(1)
// 	input <- initialValue
// 	close(input)
// 	wg.Wait()
// 	return output
// }

// func (r *recordCreator) getRecordSourcesOld(collectorSpec *databot.RecordCollectorSpec) []string {
// 	docReader := r.docReaderFn(collectorSpec.SourceURI)
// 	source, err := docReader.ReadAsString()
// 	if err != nil {
// 		logger.Errorf("unable to retrieve content from URI %s", collectorSpec.SourceURI)
// 	}

// 	result := []string{}
// 	result = append(result, source)

// 	for key, val := range collectorSpec.Params {
// 		switch key {
// 		case "fetch":
// 			result = fetch(result, val, collectorSpec)
// 		case "css:remove":
// 			result = cssRemove(result, val)
// 		case "css:select":
// 			result = cssSelect(result, val)
// 		case "css:selectEach":
// 			result = cssSelectEach(result, val)
// 		case "regexp:remove":
// 			result = regexpRemove(result, val)
// 		case "regexp:select":
// 			result = regexpSelect(result, val)
// 		case "regexp:selectEach":
// 			result = regexpSelectEach(result, val)
// 		}
// 	}
// 	return result
// }

// func resolveRelativePath(sourceURL string, relativePath string) string {
// 	if strings.HasPrefix(relativePath, "http") {
// 		return relativePath
// 	} else if strings.HasPrefix(relativePath, "/") {
// 		baseURL := regexp.MustCompile("^.+?[^/:]([?/]|$)").FindString(sourceURL)
// 		baseURL = strings.TrimRight(baseURL, "/") // remove tailing / if present
// 		return baseURL + relativePath
// 	} else if strings.HasSuffix(sourceURL, "/") {
// 		return sourceURL + relativePath
// 	} else {
// 		return sourceURL + "/" + relativePath
// 	}
// }

// func cssRemove(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	for i, block := range input {
// 		doc := html.NewDocument(block)
// 		doc.Remove(selectors...)
// 		input[i] = doc.HTML()
// 	}
// 	return input
// }

// func cssSelect(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	for i, block := range input {
// 		doc := html.NewDocument(block)
// 		doc.Select(selectors...)
// 		input[i] = doc.HTML()
// 	}
// 	return input
// }

// func cssSelectEach(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	newResult := []string{}
// 	for _, block := range input {
// 		doc := html.NewDocument(block)
// 		newResult = append(newResult, doc.HTMLEach(selectors...)...)
// 	}
// 	return newResult
// }

// func regexpRemove(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	for i, block := range input {
// 		for _, selector := range selectors {
// 			matches := regexpMatchAll(block, selector)
// 			for _, match := range matches {
// 				block = strings.Replace(block, match, "", -1)
// 			}
// 		}
// 		input[i] = block
// 	}
// 	return input
// }

// func regexpSelect(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	for i, block := range input {
// 		for _, selector := range selectors {
// 			matches := regexpMatchAll(block, selector)
// 			block = strings.Join(matches, "")
// 		}
// 		input[i] = block
// 	}
// 	return input
// }

// func regexpSelectEach(input []string, param interface{}) []string {
// 	selectors, ok := param.([]string)
// 	if !ok {
// 		panic("selector must be specified using slice of string")
// 	}
// 	newResult := []string{}
// 	for _, block := range input {
// 		for _, selector := range selectors {
// 			matches := regexpMatchAll(block, selector)
// 			newResult = append(newResult, matches...)
// 		}
// 	}
// 	return newResult
// }

// func fetch(input []string, param interface{}, spec *databot.RecordCollectorSpec) []string {
// 	if param != nil && param.(bool) {
// 		result := []string{}
// 		for _, url := range input {
// 			url = resolveRelativePath(spec.SourceURI, url)
// 			docReader := html.NewCachedDocumentReader(url, cache.Current())
// 			htmlStr, err := docReader.ReadAsString()
// 			if err != nil {
// 				logger.Errorf("unable to get html from url: %s, skipping it", url)
// 			} else {
// 				result = append(result, htmlStr)
// 			}
// 		}
// 		return result
// 	}
// 	return input
// }

// func createField(source string, spec *databot.FieldSpec) interface{} {
// 	collected := collectField(source, spec.CollectorSpec)
// 	return transformField(collected, spec.TransformerSpecs)
// }

// func collectField(source string, spec *databot.FieldCollectorSpec) interface{} {
// 	collectorType := spec.Type

// 	// no html specific collector yet
// 	// if collector := collectorMap[collectorType]; collector != nil {
// 	// 	return collector(source, spec.Params)
// 	// }

// 	if sharedCollector := fldcollector.CollectorMap[collectorType]; sharedCollector != nil {
// 		return sharedCollector(source, spec.Params)
// 	}

// 	logger.Errorf("specified collector %s is not found", collectorType)
// 	return nil
// }

// func transformField(value interface{}, specs []*databot.FieldTransformerSpec) interface{} {
// 	for _, spec := range specs {

// 		// no exclusive to html source transformer yet
// 		// if transformerFn := transformersMap[spec.Type]; transformerFn != nil {
// 		// 	value = transformerFn(value, spec.Params)
// 		// 	continue
// 		// }

// 		if sharedTransformerFn := fldtransformer.TransformersMap[spec.Type]; sharedTransformerFn != nil {
// 			value = sharedTransformerFn(value, spec.Params)
// 			continue
// 		}

// 		logger.Errorf("transformer %s not found", spec.Type)
// 	}
// 	return value
// }
