package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return isPalindromeArray(arr)
}

func isPalindromeArray(arr []int) bool {
	n := len(arr)
	for index := 0; index < n/2; index++ {
		if arr[index] != arr[n-1-index] {
			return false
		}
	}
	return true
}
