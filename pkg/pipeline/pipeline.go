package pipeline

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/processor"
)

// GetProcessors returns Processors for the input ProcessorSpec in the same order
func GetProcessors(processorSpecs []*databot.ProcessorSpec) []processor.Processor {
	processors := []processor.Processor{}
	for _, spec := range processorSpecs {
		builder := processor.GetProcessorBuilder(spec.Name)
		processors = append(processors, builder(spec.Params))
	}
	return processors
}

// CreateSingleUsePipeline creates single use pipeline with given input channel and Processors and returns output channel
func CreateSingleUsePipeline(input <-chan interface{}, processors []processor.Processor) <-chan interface{} {
	var output chan interface{}
	for _, p := range processors {
		output = make(chan interface{})
		go func(p processor.Processor, in <-chan interface{}, out chan interface{}) {
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("Error when running pipeline, returning early. Error %s", r.(error).Error())
					close(out)
				}
			}()
			p(in, out)
			close(out)
		}(p, input, output)
		input = output
	}
	return output
}
