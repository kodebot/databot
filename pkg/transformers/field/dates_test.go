package field

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	timePointer := time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC)
	tests := []TransformerTest{
		{time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC), "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
		{&timePointer, "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
		{nil, nil, nil},
		{"not a date", "not a date", nil}}

	for _, test := range tests {
		actual := formatDate(test.input, test.params)
		if test.expected != actual {
			fail(t, "formatDate not working", test.expected, actual)
		}
	}
}

func TestParseDate(t *testing.T) {
	tests := []TransformerTest{
		// {nil, nil, nil},
		// {"06/10/19", "2019-06-10 00:00:00 +0000 UTC", map[string]interface{}{
		// 	"layout": "01/02/06"}},
		// {"07/10/19", "07/10/19", nil},
		// {"2019-05-04T15:04:05Z07:00", "2019-05-04T15:04:05Z07:00", nil},
		// {"2019-05-04T15:04:05Z07:00", "2019-05-04T15:04:05Z07:00", map[string]interface{}{}},
		{"Tue, 11 Jun 2019 23:26:00 +0530", "2019-06-11 23:26:00 +0530 +0530", map[string]interface{}{"layout": "Mon, 02 Jan 2006 15:04:05 -0700"}},
		{nil, nil, nil},
		{"not a date", "not a date", nil}}

	for _, test := range tests {
		actual := parseDate(test.input, test.params)
		if !compareDateStr(test.expected, actual) {
			fail(t, "parseDate not working", test.expected, actual)
		}
	}
}

func TestUtcNow(t *testing.T) {
	tests := []struct {
		input  interface{}
		params map[string]interface{}
	}{
		{nil, nil},
		{"06/10/19", map[string]interface{}{}},
		{"not a date", map[string]interface{}{}},
		{"not a date", nil}}

	for _, test := range tests {
		actual := utcNow(test.input, test.params)
		if valTime, ok := actual.(time.Time); ok {
			now := time.Now()
			if !((valTime.Before(now) || valTime.Equal(now)) && valTime.After(now.Add(-2*time.Second))) {
				fail(t, "utcNow is not returning current time in UTC", "<<UTC_NOW>>", valTime.String())
			}
		} else {
			fail(t, "utcNow is not returning time", nil, nil)
		}
	}
}
