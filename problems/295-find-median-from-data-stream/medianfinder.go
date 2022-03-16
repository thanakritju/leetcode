package findmedianfromdatastream

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	MaxHeap MaxHeap
	MinHeap MinHeap
}

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MaxHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int {
	n := len(h)
	if n == 0 {
		return 0
	} else {
		return h[n-1]
	}
}
func (h MaxHeap) Top() int {
	n := len(h)
	if n == 0 {
		return -1
	} else {
		return h[n-1]
	}
}

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

func Constructor() MedianFinder {
	min := &MinHeap{}
	max := &MaxHeap{}
	heap.Init(min)
	heap.Init(max)
	return MedianFinder{MinHeap: *min, MaxHeap: *max}
}

func (m *MedianFinder) AddNum(num int) {
	if float64(num) > m.FindMedian() {
		heap.Push(&m.MinHeap, num)
	} else {
		heap.Push(&m.MaxHeap, num)
	}
	maxHeapLen := m.MaxHeap.Len()
	minHeapLen := m.MinHeap.Len()
	if maxHeapLen-minHeapLen > 1 {
		heap.Push(&m.MinHeap, m.MaxHeap.Pop())
	}
	if minHeapLen-maxHeapLen > 1 {
		heap.Push(&m.MaxHeap, m.MinHeap.Pop())
	}
	fmt.Printf("minHeap: %v\nmaxHeap: %v\n", m.MinHeap, m.MaxHeap)
}

func (m *MedianFinder) FindMedian() float64 {
	maxHeapLen := m.MaxHeap.Len()
	minHeapLen := m.MinHeap.Len()
	if maxHeapLen == minHeapLen {
		return float64(m.MaxHeap.Top()*m.MinHeap.Top()) / float64(2)
	}
	if maxHeapLen > minHeapLen {
		return float64(m.MaxHeap.Top())
	}
	return float64(m.MinHeap.Top())
}
