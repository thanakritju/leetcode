package recursion

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}
	reverseString(s[1:])
	temp := s[0]
	for index, char := range s[1:] {
		s[index] = char
	}
	s[len(s)-1] = temp
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	swapPairsHelper(head, true)
	return head
}

func swapPairsHelper(head *ListNode, shouldSwap bool) {
	if head == nil || head.Next == nil {
		return
	}
	if shouldSwap {
		temp := head.Next.Val
		head.Next.Val = head.Val
		head.Val = temp
	}
	swapPairsHelper(head.Next, !shouldSwap)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rest := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

func getRow(rowIndex int) []int {
	arrs := [][]int{}
	for i := 0; i < rowIndex+1; i++ {
		arr := []int{}
		for j := 0; j <= i; j++ {
			arr = append(arr, getRowHelper(i, j, arrs))
		}
		arrs = append(arrs, arr)
	}
	return arrs[rowIndex]
}

func getRowHelper(rowIndex int, columnIndex int, memory [][]int) int {
	if columnIndex == rowIndex || columnIndex == 0 {
		return 1
	}
	if len(memory) > rowIndex && len(memory[rowIndex]) > 0 {
		return memory[rowIndex][columnIndex]
	}
	return getRowHelper(rowIndex-1, columnIndex-1, memory) + getRowHelper(rowIndex-1, columnIndex, memory)
}

func fib(n int) int {
	cache := make(map[int]int)
	return fibRecursive(n, cache)
}

func fibRecursive(n int, cache map[int]int) int {
	i, ok := cache[n]
	var result int
	if ok {
		result = i
	}
	if n < 2 {
		result = n
	} else {
		result = fibRecursive(n-1, cache) + fibRecursive(n-2, cache)
	}
	cache[n] = result
	return result
}

func climbStairs(n int) int {
	arr := []int{0, 1, 2, 3}
	for {
		lenArr := len(arr)
		if lenArr > n {
			break
		}
		arr = append(arr, arr[lenArr-1]+arr[lenArr-2])
	}
	return arr[n]
}
