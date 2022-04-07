package coinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	var tests = []struct {
		input    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1}, 2, 2},
		{[]int{3, 7, 405, 436}, 8839, 25},
		{[]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864, 24},
	}

	for _, test := range tests {
		actual := coinChange(test.input, test.amount)
		assert.Equal(t, test.expected, actual, test)
	}
}
