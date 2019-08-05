package record

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/pipeline"
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
	processors := pipeline.GetProcessors(spec.PreprocessorSpecs)
	in := make(chan interface{})
	out := pipeline.CreateSingleUsePipeline(in, processors)

	go func() {
		in <- spec.SourceURI
		close(in)
	}()
	records := <-out
	collected := collect(records.([]interface{}), spec)
	return collected
}

func collect(sources []interface{}, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}
	fieldProcessorPipelineMap := make(map[string][]processor.Processor)

	for _, fieldSpec := range spec.FieldSpecs {
		processors := pipeline.GetProcessors(fieldSpec.ProcessorSpecs)
		fieldProcessorPipelineMap[fieldSpec.Name] = processors
	}

	for _, item := range sources {
		rec := make(map[string]interface{})
		for key, processors := range fieldProcessorPipelineMap {
			in := make(chan interface{})
			out := pipeline.CreateSingleUsePipeline(in, processors)
			go func() {
				in <- item
				close(in)
			}()

			res := <-out
			rec[key] = res
		}
		recs = append(recs, rec)
	}
	return recs
}
