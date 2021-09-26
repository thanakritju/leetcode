package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	var tests = []struct {
		inputArray []byte
		expected   []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	}

	for _, test := range tests {
		reverseString(test.inputArray)
		assert.Equal(t, test.expected, test.inputArray, test)
	}
}

func TestSwapPairs(t *testing.T) {
	var tests = []struct {
		inputArray []int
		expected   []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, test := range tests {
		head := swapPairs(arrayToLinklist(test.inputArray))
		assert.Equal(t, test.expected, linklistToArray(head), test)
	}
}

func TestReverseList(t *testing.T) {
	var tests = []struct {
		inputArray []int
		expected   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{2, 1}},
	}

	for _, test := range tests {
		head := reverseList(arrayToLinklist(test.inputArray))
		assert.Equal(t, test.expected, linklistToArray(head), test)
	}
}

func TestGetRow(t *testing.T) {
	var tests = []struct {
		input    int
		expected []int
	}{
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{1, []int{1, 1}},
		{0, []int{1}},
	}

	for _, test := range tests {
		actual := getRow(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
func TestFib(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{3, 2},
		{4, 3},
		{1, 1},
		{0, 0},
	}

	for _, test := range tests {
		actual := fib(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{5, 8},
		{4, 5},
		{3, 3},
		{2, 2},
		{1, 1},
	}

	for _, test := range tests {
		actual := climbStairs(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestMyPow(t *testing.T) {
	var tests = []struct {
		inputX   float64
		inputN   int
		expected float64
	}{
		{2.0, 10, 1024.0},
		{2.1, 3, 9.261},
		{2.0, -2, 0.25},
		{2.0, 0, 1.0},
		{0.00001, 2147483647, 0.0001},
	}

	for _, test := range tests {
		actual := myPow(test.inputX, test.inputN)
		assert.InDelta(t, test.expected, actual, 0.0001, test)
	}
}

func arrayToLinklist(arr []int) *ListNode {
	tempHead := &ListNode{}
	head := tempHead
	for index, val := range arr {
		head.Val = val
		if index == len(arr)-1 {
			break
		}
		head.Next = &ListNode{}
		head = head.Next
	}
	if len(arr) == 0 {
		return nil
	}
	return tempHead
}

func linklistToArray(head *ListNode) []int {
	arr := []int{}
	for {
		if head == nil {
			break
		}
		arr = append(arr, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return arr
}
