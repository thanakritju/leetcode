package relativeranks

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findRelativeRanks(score []int) []string {
	max := &MaxHeap{}
	heap.Init(max)
	for _, each := range score {
		heap.Push(max, each)
	}
	m := make(map[int]string)
	i := 1
	for len(*max) != 0 {
		val := heap.Pop(max)
		m[val.(int)] = toString(i)
		i++
	}
	ans := []string{}
	for _, each := range score {
		ans = append(ans, m[each])
	}
	return ans
}

func toString(in int) string {
	if in == 1 {
		return "Gold Medal"
	} else if in == 2 {
		return "Silver Medal"
	} else if in == 3 {
		return "Bronze Medal"
	}
	return fmt.Sprint(in)
}
