package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, test := range tests {
		output := easy.LengthOfLastWord(test.s)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("LengthOfLastWord(%s) actual: %d - expected: %d",
				test.s, output, test.expected)
		}
	}
}
