package valid_parentheses

import (
	"reflect"
	"testing"
)

var tests = []struct {
	s        string
	expected bool
}{
	{
		s:        "()",
		expected: true,
	},
	{
		s:        "()[]{}",
		expected: true,
	},
	{
		s:        "(]",
		expected: false,
	},
}

func Test_isValid(t *testing.T) {
	for _, test := range tests {
		actual := isValid(test.s)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("isValid(%s): expected %t, actual %t", test.s, test.expected, actual)
		}
	}
}
