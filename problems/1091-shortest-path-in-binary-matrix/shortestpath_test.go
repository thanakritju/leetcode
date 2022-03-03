package shortestpath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	t.Skip("come back later")
	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
	}

	for _, test := range tests {
		actual := shortestPathBinaryMatrix(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
