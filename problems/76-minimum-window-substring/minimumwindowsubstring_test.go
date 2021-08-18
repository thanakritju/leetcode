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
		{"ab", "a", "a"},
		{"ab", "A", ""},
		{"a", "b", ""},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "ba", "ab"},
		{"ZZADOBECODEBANCZZ", "ABC", "BANC"},
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
		{"cccabcccc", "ba", "ab"},
	}

	for _, test := range tests {
		actual := minWindow(test.inputS, test.inputT)
		assert.Equal(t, test.expected, actual, test)
	}
}
