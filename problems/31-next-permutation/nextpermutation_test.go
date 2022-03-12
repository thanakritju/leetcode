package nextpermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}},
		{[]int{1, 2, 4, 3}, []int{1, 3, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
	}

	for _, test := range tests {
		nextPermutation(test.input)
		assert.Equal(t, test.expected, test.input, test)
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, test := range tests {
		reverse(test.input)
		assert.Equal(t, test.expected, test.input, test)
	}
}
