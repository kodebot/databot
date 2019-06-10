// +build integration

package datafeed

import "testing"

func TestIntegrationReadFeedInfo(t *testing.T) {
	feedInfo := readFeedInfo("./test_data/only_for_reader/feed_parsing_config1.toml")

	expectedSource := "http://rss.dinamalar.com/?cat=ara1"
	if actualSource := string(feedInfo.Source); actualSource != expectedSource {
		t.Fatalf("record Source does not match. Expected: %s ** Actual: %s", expectedSource, actualSource)
	}

	expectedSourceType := "rss/atom"
	if actualSourceType := string(feedInfo.SourceType); actualSourceType != expectedSourceType {
		t.Fatalf("record sourceType does not match. Expected: %s ** Actual: %s", expectedSourceType, actualSourceType)
	}

	expectedSourceName := "dinamalar"
	if actualSourceName := string(feedInfo.SourceName); actualSourceName != expectedSourceName {
		t.Fatalf("record SourceName does not match. Expected: %s ** Actual: %s", expectedSourceName, actualSourceName)
	}

	expectedCategory := "politics"
	if actualCategory := string(feedInfo.Category); actualCategory != expectedCategory {
		t.Fatalf("record Category does not match. Expected: %s ** Actual: %s", expectedCategory, actualCategory)
	}

	expectedSchedule := "every 5 minutes"
	if actualSchedule := string(feedInfo.Schedule); actualSchedule != expectedSchedule {
		t.Fatalf("record Schedule does not match. Expected: %s ** Actual: %s", expectedSchedule, actualSchedule)
	}

	record := feedInfo.Record

	if fieldCount := len(record.Fields); fieldCount != 6 {
		t.Fatalf("expecting 6 field settings but found %d", fieldCount)
	}

	lastField := record.Fields[5]

	expectedName := "ThumbImageUrl"
	if actualName := lastField.Name; actualName != expectedName {
		t.Fatalf("field name does not match. Expected: %s ** Actual: %s", expectedName, actualName)
	}

	expectedCollectorType := "regexp"
	if actualCollectorType := string(lastField.CollectorInfo.Type); actualCollectorType != expectedCollectorType {
		t.Fatalf("field collector type does not match. Expected: %s ** Actual: %s", expectedCollectorType, actualCollectorType)
	}

	expectedCollectorTypeParamSource := "Description"
	if actualCollectorTypeParamSource := lastField.CollectorInfo.Parameters["source"].(string); actualCollectorTypeParamSource != expectedCollectorTypeParamSource {
		t.Fatalf("field collector type parameter (Source) does not match. Expected: %s ** Actual: %s", expectedCollectorTypeParamSource, actualCollectorTypeParamSource)
	}
}
