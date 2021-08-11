package isgraphbipartite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected bool
	}{
		{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
		{[][]int{{1}, {0}, {4}, {4}, {2, 3}}, true},
	}

	for _, test := range tests {
		actual := isBipartite(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
