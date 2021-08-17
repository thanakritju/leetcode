package findanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	var tests = []struct {
		inputS   string
		inputP   string
		expected []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"ccc", "ab", []int{}},
		{"ccc", "abss", []int{}},
	}

	for _, test := range tests {
		actual := findAnagrams(test.inputS, test.inputP)
		assert.Equal(t, test.expected, actual, test)
	}
}
