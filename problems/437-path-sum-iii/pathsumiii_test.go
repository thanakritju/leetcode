package pathsumiii

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	var tests = []struct {
		inputArray     []string
		inputTargetSum int
		expected       int
	}{
		{[]string{"10", "5", "-3", "3", "2", "null", "11", "3", "-2", "null", "1"}, 8, 3},
		{[]string{"10", "5", "-2", "2", "-1", "null", "9", "null", "7", "null", "3"}, 7, 4},
		{[]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"}, 22, 3},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := pathSum(root, test.inputTargetSum)
		assert.Equal(t, test.expected, actual, test)
	}
}

func insertLevelOrder(arr []string, root *TreeNode, i int, n int) *TreeNode {
	if i < n {
		value, err := strconv.Atoi(arr[i])
		if err != nil {
			return root
		}
		temp := TreeNode{Val: value}
		root = &temp
		root.Left = insertLevelOrder(arr, root.Left, 2*i+1, n)
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2, n)
	}
	return root
}
