package minimumwindowsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	var tests = []struct {
		inputS   string
		inputT   string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "ba", "ab"},
		{"cccabcccc", "ba", "ab"},
	}

	for _, test := range tests {
		actual := minWindow(test.inputS, test.inputT)
		assert.Equal(t, test.expected, actual, test)
	}
}
