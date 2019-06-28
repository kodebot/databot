package field

import (
	"testing"

	"github.com/kodebot/databot/pkg/databot"
)

func TestTransformTransformerFound(t *testing.T) {

	value := " test string "

	specs := []databot.FieldTransformerSpec{{
		Type: Trim}}

	actual := Transform(value, specs)
	expected := "test string"
	if actual != expected {
		fail(t, "transform not applied", expected, actual.(string))
	}
}

func TestTransformUnknownTransformer(t *testing.T) {
	value := " test string "

	specs := []databot.FieldTransformerSpec{{
		Type: "Unknown"}}

	actual := Transform(value, specs)
	expected := " test string "
	if actual != expected {
		fail(t, "transforms should not be applied but appears to be applied", expected, actual.(string))
	}
}

func TestTransformMultipleTransformers(t *testing.T) {
	value := " test string "

	specs := []databot.FieldTransformerSpec{
		{
			Type: TrimLeft},
		{
			Type: TrimRight}}

	actual := Transform(value, specs)
	expected := "test string"
	if actual != expected {
		fail(t, "not all the transforms appear to be applied", expected, actual.(string))
	}
}

func TestTransformMultipleTransformersAndUnknown(t *testing.T) {
	value := " test string "

	specs := []databot.FieldTransformerSpec{
		{
			Type: "Unknown"},
		{
			Type: TrimRight}}

	actual := Transform(value, specs)

	expected := " test string"
	if actual != expected {
		fail(t, "valid transformers not appear to be applied", expected, actual.(string))
	}
}
