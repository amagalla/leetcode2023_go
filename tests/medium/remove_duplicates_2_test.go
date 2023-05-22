package medium

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/medium"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		nums       []int
		intExp     int
		numsArrExp []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3, 3, 3}},
	}

	for _, test := range tests {
		integer, newArr := medium.RemoveDuplicates2(test.nums)
		if integer != test.intExp || !reflect.DeepEqual(newArr, test.numsArrExp) {
			t.Errorf("RemoveDuplicates2(%v) actual: %d, %v - expected: %d, %v",
				test.nums, integer, newArr, test.intExp, test.numsArrExp)
		}
	}
}
