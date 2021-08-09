package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   bool
	}{
		{[]string{"1", "2", "2", "null", "3", "null", "3"}, false},
		{[]string{"1", "2", "2", "3", "4", "4", "3"}, true},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := isSymmetric(root)
		assert.Equal(t, test.expected, actual, test)
	}
}
