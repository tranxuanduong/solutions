package binary_search

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums     []int
	target   int
	expected int
}{
	{
		nums:     []int{-1, 0, 3, 5, 9, 12},
		target:   9,
		expected: 4,
	},
	{
		nums:     []int{-1, 0, 3, 5, 9, 12},
		target:   2,
		expected: -1,
	},
}

func Test_search(t *testing.T) {
	for _, test := range tests {
		actual := search(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("search(%+v, %d): expected %+v, actual %+v", test.nums, test.target, test.expected, actual)
		}
	}
}
