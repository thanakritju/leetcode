package trappingrainwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, test := range tests {
		assert.Equal(t, trap(test.input), test.expected)
	}
}

func TestGetWaterBetweenIndexes(t *testing.T) {
	var tests = []struct {
		inputIndex1 int
		inputIndex2 int
		inputArray  []int
		expected    int
	}{
		{3, 7, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 4},
	}

	for _, test := range tests {
		assert.Equal(t, getWaterBetweenIndexes(test.inputIndex1, test.inputIndex2, test.inputArray), test.expected)
	}
}
