package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		output := easy.RomanToInteger(test.input)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("RomanToInteger(%s) actual: %d - expected: %d",
				test.input, output, test.expected)
		}
	}
}
