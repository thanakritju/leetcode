package longestcommonsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	var tests = []struct {
		t1       string
		t2       string
		expected int
	}{
		{"a", "a", 1},
		{"a", "b", 0},
		{"ace", "abcde", 3},
		{"upr", "urp", 2},
		{"ezupkr", "ubmrapg", 2},
	}

	for _, test := range tests {
		actual := longestCommonSubsequence(test.t1, test.t2)
		assert.Equal(t, test.expected, actual, test)
	}
}
