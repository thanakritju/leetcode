package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		inputArray []int
		target     int
		expected   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, test := range tests {
		actual := search(test.inputArray, test.target)
		assert.Equal(t, test.expected, actual, test)
	}
}
