package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums       []int
		val        int
		integerExp int
		numsArrExp []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2, 2, 3}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4, 0, 4, 2}},
	}

	for _, test := range tests {
		integer, newArr := easy.RemoveElement(test.nums, test.val)

		if integer != test.integerExp || !reflect.DeepEqual(newArr, test.numsArrExp) {
			t.Errorf("RemoveElement(%v, %d) actual: %d %v - expected %d %v",
				test.nums, test.val, integer, newArr, test.integerExp, test.numsArrExp)
		}
	}
}
