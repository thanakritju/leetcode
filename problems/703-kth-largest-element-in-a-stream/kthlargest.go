package kthlargest

import (
	"container/heap"
)

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MaxHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	MaxHeap MaxHeap
	MinHeap MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	max := &MaxHeap{}
	min := &MinHeap{}
	heap.Init(max)
	heap.Init(min)
	for _, each := range nums {
		heap.Push(max, each)
	}

	i := 0
	for i < k && len(*max) != 0 {
		heap.Push(min, heap.Pop(max))
		i++
	}
	return KthLargest{MaxHeap: *max, MinHeap: *min, k: k}
}

func (k *KthLargest) Add(val int) int {
	if len(k.MinHeap) == 0 || val > k.MinHeap[0] {
		heap.Push(&k.MinHeap, val)
	} else {
		heap.Push(&k.MaxHeap, val)
	}
	for len(k.MinHeap) < k.k && len(k.MaxHeap) != 0 {
		heap.Push(&k.MinHeap, heap.Pop(&k.MaxHeap))
	}
	if len(k.MinHeap) > k.k {
		heap.Push(&k.MaxHeap, heap.Pop(&k.MinHeap))
	}
	return k.MinHeap[0]
}
