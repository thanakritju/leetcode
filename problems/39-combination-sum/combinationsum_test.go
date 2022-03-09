package combinationsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	var tests = []struct {
		inputArray  []int
		inputTarget int
		expected    [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{2, 7, 6, 3, 5, 1}, 9, [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 2, 3}, {1, 1, 1, 1, 5}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 3, 3}, {1, 1, 1, 6}, {1, 1, 2, 2, 3}, {1, 1, 2, 5}, {1, 1, 7}, {1, 2, 2, 2, 2}, {1, 2, 3, 3}, {1, 2, 6}, {1, 3, 5}, {2, 2, 2, 3}, {2, 2, 5}, {2, 7}, {3, 3, 3}, {3, 6}}},
	}

	for _, test := range tests {
		actual := combinationSum(test.inputArray, test.inputTarget)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestArePermutation(t *testing.T) {
	var tests = []struct {
		inputArray  []int
		inputTarget []int
		expected    bool
	}{
		{[]int{2, 1, 1, 2, 1, 1, 1}, []int{2, 1, 2, 1, 1, 1, 1}, true},
	}

	for _, test := range tests {
		actual := arePermutation(test.inputArray, test.inputTarget)
		assert.Equal(t, test.expected, actual, test)
	}
}
