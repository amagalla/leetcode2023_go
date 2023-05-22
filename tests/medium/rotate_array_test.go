package medium

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/medium"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, test := range tests {
		output := medium.Rotate(test.nums, test.k)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Rotate(%v, %d) actual: %v - expected: %v",
				test.nums, test.k, output, test.expected)
		}
	}
}
