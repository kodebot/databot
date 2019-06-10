package collectors

import "testing"

var regexpTests = []CollectorTest{
	{nil, nil, nil},
	{"no regex", nil, nil},
	{"regex without data group", nil, map[string]interface{}{"expr": ".*result.*"}},
	{"invalid regex", nil, map[string]interface{}{"expr": "??..*result.*"}},
	{"get result from this", "result", map[string]interface{}{"expr": ".*(?P<data>result).*"}},
	{nil, nil, map[string]interface{}{"expr": ".*(?P<data>result).*"}}}

func TestRegexp(t *testing.T) {
	for _, test := range regexpTests {
		actual := regex(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "regexp collector not working", test.expected, actual)
		}
	}
}
