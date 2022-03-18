package findmedianfromdatastream

import (
	"container/heap"
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
		heap.Push(&m.MinHeap, heap.Pop(&m.MaxHeap))
	}
	if minHeapLen-maxHeapLen > 1 {
		heap.Push(&m.MaxHeap, heap.Pop(&m.MinHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	maxHeapLen := m.MaxHeap.Len()
	minHeapLen := m.MinHeap.Len()
	if maxHeapLen == 0 && minHeapLen == 0 {
		return 0
	}
	if maxHeapLen > minHeapLen {
		return float64(m.MaxHeap[0])
	}
	if maxHeapLen < minHeapLen {
		return float64(m.MinHeap[0])
	}
	return float64(m.MaxHeap[0]+m.MinHeap[0]) / float64(2)
}
