package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range tests {
		output := easy.MaxProfit(test.prices)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("MaxProfit(%v) actual: %d - expected: %d",
				test.prices, output, test.expected)
		}
	}
}
