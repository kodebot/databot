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

	// go func() {
	// 	for control := range output.Control {
	// 		// drain
	// 		fmt.Printf("%+v", control)
	// 	}
	// }()

	records := <-output.Data
	fmt.Printf("%+v", records.([]interface{}))

	collected := collect(records.([]interface{}), spec)
	return collected
}

func collect(sources []interface{}, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}

	fieldProcessorPipelineMap := make(map[string][]processor.Flow)

	for _, fieldSpec := range spec.FieldSpecs {
		input, output := buildProcessorPipeline(fieldSpec.ProcessorSpecs)
		fieldProcessorPipelineMap[fieldSpec.Name] = []processor.Flow{input, output}
	}

	for _, item := range sources {
		rec := make(map[string]interface{})
		for key, val := range fieldProcessorPipelineMap {
			in := val[0]
			out := val[1]

			go func() {
				in.Data <- item
			}()

			go func() {
				for control := range out.Control {
					// drain
					fmt.Printf("%+v", control)
				}
			}()

			rec[key] = <-out.Data
		}
		recs = append(recs, rec)
	}
	fmt.Printf("%+v", sources)
	return recs
}

// https://whiskybadger.io/post/introducing-go-pipeline/
func buildProcessorPipeline(processorSpecs []*databot.ProcessorSpec) (processor.Flow, processor.Flow) {
	data := make(chan interface{})
	control := make(chan processor.ControlMessage)

	input := processor.Flow{
		Data:    data,
		Control: control,
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
