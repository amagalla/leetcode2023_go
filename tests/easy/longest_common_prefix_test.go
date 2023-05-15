package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		output := easy.LongestCommonPrefix(test.strs)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("LongestCommonPrefix(%s) actual: %s - exptected: %s",
				test.strs, output, test.expected)
		}
	}
}
