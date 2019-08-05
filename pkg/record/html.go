package record

import (
	"fmt"

	"github.com/kodebot/databot/pkg/pipeline"

	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
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
	operators := buildProcessorPipeline(spec.PreprocessorSpecs)

	in := make(chan interface{})

	out := make(chan interface{})
	tempIn := in
	for _, operator := range operators {
		go func(op pipeline.Operator, in chan interface{}, o chan interface{}) {
			op(in, o)
			close(o)
		}(operator, tempIn, out)
		tempIn = out
		out = make(chan interface{})
	}
	go func() {
		in <- spec.SourceURI
		close(in)
	}()

	records := <-tempIn

	fmt.Printf("%+v", records.([]interface{}))

	collected := collect(records.([]interface{}), spec)
	return collected
}

func collect(sources []interface{}, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}

	fieldProcessorPipelineMap := make(map[string][]pipeline.Operator)

	for _, fieldSpec := range spec.FieldSpecs {
		operators := buildProcessorPipeline(fieldSpec.ProcessorSpecs)
		fieldProcessorPipelineMap[fieldSpec.Name] = operators
	}

	for _, item := range sources {
		rec := make(map[string]interface{})
		for key, val := range fieldProcessorPipelineMap {

			in := make(chan interface{})

			out := make(chan interface{})
			tempIn := in
			for _, op := range val {
				go func(op1 pipeline.Operator, in chan interface{}, o chan interface{}) {
					op1(in, o)
					close(o)
				}(op, tempIn, out)
				tempIn = out
				out = make(chan interface{})
			}

			go func() {
				in <- item
				close(in)
			}()

			res := <-tempIn
			rec[key] = res

		}
		recs = append(recs, rec)
	}
	fmt.Printf("%+v", sources)
	return recs
}

func buildProcessorPipeline(processorSpecs []*databot.ProcessorSpec) []pipeline.Operator {
	var operators []pipeline.Operator
	for _, spec := range processorSpecs {
		builder := pipeline.Get(spec.Name)
		operators = append(operators, builder(spec.Params))
	}
	return operators

}
