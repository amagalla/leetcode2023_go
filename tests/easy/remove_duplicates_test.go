package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums       []int
		integerExp int
		numsArrExp []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}},
	}

	for _, test := range tests {
		intOutput, numsOutput := easy.RemoveDuplicates(test.nums)

		if intOutput != test.integerExp || !reflect.DeepEqual(numsOutput, test.numsArrExp) {
			t.Errorf("RemoveDuplicates(%v) actual: %v %d - expected %v, %d",
				test.nums, intOutput, numsOutput, test.integerExp, test.numsArrExp)
		}
	}
}
