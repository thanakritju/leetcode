package topkfrequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	var tests = []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 8, 8, 8, 8, 8, 3}, 2, []int{8, 1}},
		{[]int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		actual := topKFrequent(test.nums, test.k)
		assert.Equal(t, test.expected, actual, test)
	}
}
