package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(&nums, 0, len(nums)-1, len(nums)-k+1)
}

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func quickSelect(nums *[]int, left, right, k int) int {
	if k > 0 && k <= right-left+1 {
		index := partition(nums, left, right)

		if index-left == k-1 {
			return (*nums)[index]
		}

		if index-left > k-1 {
			return quickSelect(nums, left, index-1, k)
		}

		return quickSelect(nums, index+1, right, k-index+left-1)
	}
	return -1
}

func partition(nums *[]int, left, right int) int {
	num := (*nums)[right]
	i := left
	for j := left; j < right; j++ {
		if (*nums)[j] <= num {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i += 1
		}
	}
	(*nums)[i], (*nums)[right] = (*nums)[right], (*nums)[i]
	return i
}

type pq []int

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Pop() interface{} {
	temp := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return temp
}

func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func findKthLargest3(nums []int, k int) int {
	res := &pq{}
	for _, num := range nums {
		heap.Push(res, num)
		if res.Len() > k {
			heap.Pop(res)
		}
	}
	return heap.Pop(res).(int)
}
