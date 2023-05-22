package medium

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/medium"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, test := range tests {
		output := medium.ReverseWords(test.s)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Rotate(%s) actual: %s - expected: %s",
				test.s, output, test.expected)
		}
	}
}
