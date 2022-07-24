package searchmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	var tests = []struct {
		mat      [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{-5}}, -5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
	}

	for _, test := range tests {
		actual := searchMatrix(test.mat, test.target)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		arr      []int
		target   int
		expected bool
	}{
		{[]int{1, 4, 7, 11, 15}, 5, false},
		{[]int{1, 4}, 1, true},
		{[]int{1}, 1, true},
		{[]int{1, 4}, 4, true},
		{[]int{2, 5, 8, 12, 19}, 5, true},
		{[]int{1, 4, 7, 11, 15}, 11, true},
	}

	for _, test := range tests {
		actual := binarySearch(test.arr, test.target)
		assert.Equal(t, test.expected, actual, test)
	}

}
