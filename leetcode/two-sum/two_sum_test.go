package two_sum

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

func Test_twoSumBruceForce(t *testing.T) {
	for _, test := range tests {
		actual := twoSumBruceForce(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("twoSum(%+v, %d): expected %+v, actual %+v", test.nums, test.target, test.expected, actual)
		}
	}
}

func Test_twoSumWithMap(t *testing.T) {
	for _, test := range tests {
		actual := twoSumWithMap(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("twoSum(%+v, %d): expected %+v, actual %+v", test.nums, test.target, test.expected, actual)
		}
	}
}
