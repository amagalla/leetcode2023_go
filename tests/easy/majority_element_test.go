package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		output := easy.MajorityElement(test.nums)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("MajorityElement(%v) actual: %d - expected: %d",
				test.nums, output, test.expected)
		}
	}
}
