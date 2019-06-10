package transformers

import (
	"testing"
	"time"
)

var timePointer = time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC)
var formatDateTests = []TransformerTest{
	{time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC), "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
	{&timePointer, "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
	{nil, nil, nil},
	{"not a date", "not a date", nil}}

func TestFormatDate(t *testing.T) {
	for _, test := range formatDateTests {
		actual := formatDate(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "formatDate not working", test.expected, actual)
		}
	}
}

var parseDateTests = []TransformerTest{
	{nil, nil, nil},
	{"06/10/19", "2019-06-10 00:00:00 +0000 UTC", map[string]interface{}{
		"layout": "01/02/06"}},
	{"07/10/19", "07/10/19", nil},
	{"2019-05-04T15:04:05Z07:00", "2019-05-04T15:04:05Z07:00", nil},
	{"2019-05-04T15:04:05Z07:00", "2019-05-04T15:04:05Z07:00", map[string]interface{}{}},
	{nil, nil, nil},
	{"not a date", "not a date", nil}}

func TestParseDate(t *testing.T) {
	for _, test := range parseDateTests {
		actual := parseDate(test.input, test.parameters)
		if !compareDateStr(test.expected, actual) {
			fail(t, "parseDate not working", test.expected, actual)
		}
	}
}
