package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, test := range tests {
		output := easy.IsSubsequence(test.s, test.t)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("IsSubsequence(%s, %s) actual: %v - expected: %v",
				test.s, test.t, output, test.expected)
		}
	}
}
