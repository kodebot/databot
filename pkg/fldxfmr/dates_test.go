package fldxfmr

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	time := time.Date(2016, 10, 10, 1, 1, 1, 1, time.UTC)
	tests := []TransformerTest{
		{time, "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
		{&time, "2016-10-10 01:01:01.000000001 +0000 UTC", nil},
		{nil, nil, nil},
		{"not a date", nil, nil}}

	for _, test := range tests {
		actual := formatDate(test.input, test.params)
		if test.expected != actual {
			fail(t, "formatDate not working", test.expected, actual)
		}
	}
}

func TestParseDate(t *testing.T) {
	tests := []TransformerTest{
		{"Tue, 11 Jun 2019 23:26:00 +0530", "2019-06-11 23:26:00 +0530 +0530", map[string]interface{}{"layout": "Mon, 02 Jan 2006 15:04:05 -0700"}},
		{"06/10/19", "2019-06-10 00:00:00 +0000 UTC", map[string]interface{}{"layout": "01/02/06"}},
		{"2019-07-02T16:05:15Z", "2019-07-02 16:05:15 +0000 UTC", map[string]interface{}{}},
		{"2019-07-02T16:05:15+05:30", "2019-07-02 16:05:15 +0530 IST", map[string]interface{}{"location": "Asia/Kolkata"}},
		{"2019-07-02T16:05:15Z", "2019-07-02 16:05:15 +0000 UTC", nil},
		{"2019-07-02T16:05:15+05:30", "2019-07-02 16:05:15 +0530 +0530", map[string]interface{}{}},
		{"2019-07-02T16:05:15+05:30", "2019-07-02 16:05:15 +0530 +0530", nil},
		{"07/10/19", nil, nil},
		{"2019-05-04T15:04:05Z07:00", nil, nil},
		{"2019-05-04T15:04:05Z07:00", nil, map[string]interface{}{}},
		{"2019-05-04T15:04:05Z07:00", nil, map[string]interface{}{"location": 1}},
		{"2019-05-04T15:04:05Z07:00", nil, map[string]interface{}{"location": "foo/bar"}},
		{nil, nil, nil},
		{"not a date", nil, nil}}

	for _, test := range tests {
		actual := parseDate(test.input, test.params)
		if !compareDateStr(test.expected, actual) {
			fail(t, "parseDate not working", test.expected, actual)
		}
	}
}
