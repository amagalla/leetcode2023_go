package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestMerge(t *testing.T) {
	testing := []struct {
		nums1     []int
		m         int
		nums2     []int
		n         int
		expecting []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, test := range testing {
		output := easy.Merge(test.nums1, test.m, test.nums2, test.n)

		if !reflect.DeepEqual(output, test.expecting) {
			t.Errorf("Merge(%v, %d, %v, %d) actual: %d - expected: %d",
				test.nums1, test.m, test.nums2, test.n, output, test.expecting)
		}
	}
}
