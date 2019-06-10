package transformers

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {

	date := time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC)

	actual := formatDate(date, nil)
	expected := "2016-10-10 01:01:01.000000001 +0000 UTC"

	if expected != actual {
		t.Fatalf("date (value) is not formated using default format. EXPECTED: >>%s<<, ACTUAL: >>%s<<", expected, actual)
	}

	actual = formatDate(&date, nil)
	expected = "2016-10-10 01:01:01.000000001 +0000 UTC"

	if expected != actual {
		t.Fatalf("date (pointer) is not formated using default format. EXPECTED: >>%s<<, ACTUAL: >>%s<<", expected, actual)
	}

	actual = formatDate(nil, nil)

	if nil != actual {
		t.Fatalf("nil value causes failure. EXPECTED: >>NIL<<, ACTUAL: >>%s<<", actual)
	}

	nonDateValue := "not a date"
	actual = formatDate(nonDateValue, nil)

	if nonDateValue != actual {
		t.Fatalf("nil value causes failure. EXPECTED: >>%s<<, ACTUAL: >>%s<<", nonDateValue, actual)
	}
}
