package binarytree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   []int
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7"}, []int{1, 2, 4, 5, 3, 6, 7}},
		{[]string{"1", "null", "2", "null", "null", "3"}, []int{1, 2, 3}},
		{[]string{"1"}, []int{1}},
		{[]string{"1", "2"}, []int{1, 2}},
		{[]string{"1", "null", "2"}, []int{1, 2}},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := preorderTraversal(root)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestInorderTraversal(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   []int
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7"}, []int{4, 2, 5, 1, 6, 3, 7}},
		{[]string{"1", "null", "2", "null", "null", "3"}, []int{1, 3, 2}},
		{[]string{"1"}, []int{1}},
		{[]string{"1", "2"}, []int{2, 1}},
		{[]string{"1", "null", "2"}, []int{1, 2}},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := inorderTraversal(root)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestPostorderTraversal(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   []int
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7"}, []int{4, 5, 2, 6, 7, 3, 1}},
		{[]string{"1", "null", "2", "null", "null", "3"}, []int{3, 2, 1}},
		{[]string{"1"}, []int{1}},
		{[]string{"1", "2"}, []int{2, 1}},
		{[]string{"1", "null", "2"}, []int{2, 1}},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := postorderTraversal(root)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestLevelOrder(t *testing.T) {
	var tests = []struct {
		inputArray []string
		expected   [][]int
	}{
		{[]string{"3", "9", "20", "null", "null", "15", "7"}, [][]int{{3}, {9, 20}, {15, 7}}},
		{[]string{"1", "2", "3", "4", "5", "6", "7"}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{[]string{"1"}, [][]int{{1}}},
	}

	for _, test := range tests {
		root := insertLevelOrder(test.inputArray, &TreeNode{}, 0, len(test.inputArray))
		actual := levelOrder(root)
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
