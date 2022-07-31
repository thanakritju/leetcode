package kthlargestelementinanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	var tests = []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, test := range tests {
		actual := findKthLargest(test.arr, test.k)
		assert.Equal(t, test.expected, actual, test)
		actual = findKthLargest2(test.arr, test.k)
		assert.Equal(t, test.expected, actual, test)
	}
}
