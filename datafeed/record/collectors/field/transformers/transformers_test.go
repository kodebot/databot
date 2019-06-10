package transformers

import "testing"

func TestTransformTransformerFound(t *testing.T) {

	value := " test string "

	tranformersInfo := []TransformerInfo{{
		Transformer: Trim}}

	actual := Transform(value, tranformersInfo)
	expected := "test string"
	if actual != expected {
		fail(t, "transform not applied", expected, actual.(string))
	}
}

func TestTransformUnknownTransformer(t *testing.T) {
	value := " test string "

	tranformersInfo := []TransformerInfo{{
		Transformer: "Unknown"}}

	actual := Transform(value, tranformersInfo)
	expected := " test string "
	if actual != expected {
		fail(t, "transforms should not be applied but appears to be applied", expected, actual.(string))
	}
}

func TestTransformMultipleTransformers(t *testing.T) {
	value := " test string "

	tranformersInfo := []TransformerInfo{
		{
			Transformer: TrimLeft},
		{
			Transformer: TrimRight}}

	actual := Transform(value, tranformersInfo)
	expected := "test string"
	if actual != expected {
		fail(t, "not all the transforms appear to be applied", expected, actual.(string))
	}
}

func TestTransformMultipleTransformersAndUnknown(t *testing.T) {
	value := " test string "

	tranformersInfo := []TransformerInfo{
		{
			Transformer: "Unknown"},
		{
			Transformer: TrimRight}}

	actual := Transform(value, tranformersInfo)

	expected := " test string"
	if actual != expected {
		fail(t, "valid transformers not appear to be applied", expected, actual.(string))
	}
}

func fail(t *testing.T, message string, expected string, actual string) {
	t.Fatalf("%s. EXPECTED: >>%s<<, ACTUAL: >>%s<<", message, expected, actual)
}
