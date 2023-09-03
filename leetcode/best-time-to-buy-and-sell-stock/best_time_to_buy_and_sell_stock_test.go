package best_time_to_buy_and_sell_stock

import (
	"reflect"
	"testing"
)

var tests = []struct {
	prices   []int
	expected int
}{
	{
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	},
	{
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	},
}

func Test_maxProfitWithBruceForce(t *testing.T) {
	for _, test := range tests {
		actual := maxProfitWithBruceForce(test.prices)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("maxProfitWithBruceForce(%+v): expected %d, actual %d", test.prices, test.expected, actual)
		}
	}
}

func Test_maxProfitWithDP(t *testing.T) {
	for _, test := range tests {
		actual := maxProfitWithDP(test.prices)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("maxProfitWithDP(%+v): expected %d, actual %d", test.prices, test.expected, actual)
		}
	}
}
