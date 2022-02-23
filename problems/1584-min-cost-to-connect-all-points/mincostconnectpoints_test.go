package mincosttoconnectallpoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	t.Skip("skipping testing")
	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{0, 0}}, 0},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, 18},
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}}, 20},
		{[][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}, 4},
		{[][]int{{-1000000, -1000000}, {1000000, 1000000}}, 4000000},
	}

	for _, test := range tests {
		actual := minCostConnectPoints(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
