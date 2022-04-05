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
	}

	for _, test := range tests {
		actual := coinChange(test.input, test.amount)
		assert.Equal(t, test.expected, actual, test)
	}
}
