package easy

import (
	"reflect"
	"testing"

	"github.com/amagalla/leetcode/easy"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "but", 3},
		{"leetcode", "leeto", -1},
	}

	for _, test := range tests {
		output := easy.StrStr(test.haystack, test.needle)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("StrStr(%s, %s) actual: %d - expected %d",
				test.haystack, test.needle, output, test.expected)
		}
	}
}
