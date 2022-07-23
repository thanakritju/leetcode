package palindromelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		head     []int
		expected bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
	}

	for _, test := range tests {
		actual := isPalindrome(convertToLinkedList(test.head))
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestIsPalindromeArray(t *testing.T) {
	var tests = []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2}, false},
	}

	for _, test := range tests {
		actual := isPalindromeArray(test.arr)
		assert.Equal(t, test.expected, actual, test)
	}
}

func convertToLinkedList(input []int) *ListNode {
	head := &ListNode{}
	temp := head
	for i, num := range input {
		temp.Val = num
		if i != len(input)-1 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}
	return head
}
