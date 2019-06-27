package field

import "testing"

var regexpTests = []TransformerTest{
	{nil, nil, nil},
	{"no regex", "no regex", nil},
	{"regex without data group", "regex without data group", map[string]interface{}{"expr": ".*result.*"}},
	{"regex without data group", "", map[string]interface{}{"expr": ".*result.*", "fallbackValue": ""}},
	{"invalid regex", "invalid regex", map[string]interface{}{"expr": "??..*result.*"}},
	{"invalid regex", "fallback", map[string]interface{}{"expr": "??..*result.*", "fallbackValue": "fallback"}},
	{"get result from this", "result", map[string]interface{}{"expr": ".*(?P<data>result).*"}},
	{nil, nil, map[string]interface{}{"expr": ".*(?P<data>result).*"}}}

func TestRegexp(t *testing.T) {
	for _, test := range regexpTests {
		actual := regex(test.input, test.params)
		if test.expected != actual {
			fail(t, "regexp collector not working", test.expected, actual)
		}
	}
}
