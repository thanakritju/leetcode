package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	var tests = []struct {
		inputArray     []string
		inputTargetSum int
		expected       bool
	}{
		{[]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"}, 22, true},
		{[]string{"1", "2"}, 0, false},
		{[]string{"1", "2"}, 1, false},
		{[]string{"1", "2", "3"}, 5, false},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := hasPathSum(root, test.inputTargetSum)
		assert.Equal(t, test.expected, actual, test)
	}
}
