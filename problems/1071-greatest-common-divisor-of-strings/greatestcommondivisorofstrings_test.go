package greatestcommondivisorofstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	var tests = []struct {
		word1    string
		word2    string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}

	for _, test := range tests {
		actual := gcdOfStrings(test.word1, test.word2)
		assert.Equal(t, test.expected, actual, test)
	}
}
