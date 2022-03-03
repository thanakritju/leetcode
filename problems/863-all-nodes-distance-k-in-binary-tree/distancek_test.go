package distancek

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceK(t *testing.T) {
	var tests = []struct {
		inputArray  []string
		inputTarget int
		inputK      int
		expected    []int
	}{
		{[]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"}, 5, 2, []int{7, 4, 1}},
		{[]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "null", "5", "1"}, 5, 2, []int{11, 13, 4}},
		{[]string{"1"}, 1, 3, []int{}},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := distanceK(root, &TreeNode{Val: test.inputTarget}, test.inputK)
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
