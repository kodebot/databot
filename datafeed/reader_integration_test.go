// build integration

package datafeed

import "testing"

func TestReadRecordSettings(t *testing.T) {
	recordSettings := readRecordSettings("./test_data/feed_parsing_config1.toml")

	if count := len(recordSettings); count != 1 {
		t.Fatalf("expecting one record settings but found %d", count)
	}

	first := recordSettings[0]

	expectedType := "feed"
	if actualType := string(first.Type); actualType != expectedType {
		t.Fatalf("record setting type does not match. Expected: %s ** Actual: %s", expectedType, actualType)
	}

	expectedSource := "http://rss.dinamalar.com/?cat=ara1"
	if actualSource := string(first.Source); actualSource != expectedSource {
		t.Fatalf("record Source does not match. Expected: %s ** Actual: %s", expectedSource, actualSource)
	}

	if fieldCount := len(first.FieldSettings); fieldCount != 6 {
		t.Fatalf("expecting 6 field settings but found %d", fieldCount)
	}

	lastField := first.FieldSettings[5]

	expectedName := "ThumbImageUrl"
	if actualName := lastField.Name; actualName != expectedName {
		t.Fatalf("field name does not match. Expected: %s ** Actual: %s", expectedName, actualName)
	}

	expectedCollectorType := "regexp"
	if actualCollectorType := string(lastField.CollectorSetting.Type); actualCollectorType != expectedCollectorType {
		t.Fatalf("field collector type does not match. Expected: %s ** Actual: %s", expectedCollectorType, actualCollectorType)
	}

	expectedCollectorTypeParamSource := "Description"
	if actualCollectorTypeParamSource := lastField.CollectorSetting.Parameters["Source"].(string); actualCollectorTypeParamSource != expectedCollectorTypeParamSource {
		t.Fatalf("field collector type parameter (Source) does not match. Expected: %s ** Actual: %s", expectedCollectorTypeParamSource, actualCollectorTypeParamSource)
	}
}
