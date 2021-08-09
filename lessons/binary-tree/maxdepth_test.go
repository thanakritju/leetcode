package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   int
	}{
		{[]string{"3", "9", "20", "null", "null", "15", "7"}, 3},
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8"}, 4},
		{[]string{"0"}, 1},
		{[]string{"1", "null", "2"}, 2},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := maxDepth(root)
		assert.Equal(t, test.expected, actual, test)
	}
}
