package permuteunique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	var tests = []struct {
		inputArray []int
		expected   [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {2, 1, 1}, {1, 2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}},
	}

	for _, test := range tests {
		actual := permuteUnique(test.inputArray)
		assert.Equal(t, test.expected, actual, test)
	}
}
