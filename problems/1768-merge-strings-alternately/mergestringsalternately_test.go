package mergestringsalternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	var tests = []struct {
		word1    string
		word2    string
		expected string
	}{
		{"abc", "pqr", "apbqcr"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		actual := mergeAlternately(test.word1, test.word2)
		assert.Equal(t, test.expected, actual, test)
	}
}
