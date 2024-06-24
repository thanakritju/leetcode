package minsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinsubarray(t *testing.T) {
	var tests = []struct {
		arr      []int
		p        int
		expected int
	}{
		{[]int{3, 1, 4, 2}, 6, 1},
		{[]int{6, 3, 5, 2}, 9, 2},
		{[]int{1, 2, 3}, 3, 0},
		{[]int{1, 2, 3}, 7, -1},
	}

	for _, test := range tests {
		actual := minSubarray(test.arr, test.p)
		assert.Equal(t, test.expected, actual, test)
	}
}
