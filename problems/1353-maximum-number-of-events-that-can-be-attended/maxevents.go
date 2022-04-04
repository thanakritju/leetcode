package maxevents

import (
	"container/heap"
)

type event [][]int
type minHeap [][]int

func maxEvents(events [][]int) int {
	allevents := event(events)
	attendables := minHeap([][]int{})
	heap.Init(&attendables)
	heap.Init(&allevents)
	joined := 0
	for day := 1; day <= 10000; day++ {
		for len(allevents) > 0 && allevents[0][0] <= day {
			heap.Push(&attendables, heap.Pop(&allevents))
		}

		for len(attendables) > 0 && attendables[0][1] < day {
			heap.Pop(&attendables)
		}

		if len(attendables) > 0 {
			heap.Pop(&attendables)
			joined += 1
		} else {
			if len(allevents) == 0 {
				break
			}
		}
	}
	return joined
}

func (e event) Len() int { return len(e) }
func (e event) Less(i, j int) bool {
	return e[i][0] < e[j][0]
}
func (e event) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e *event) Push(x interface{}) {
	*e = append(*e, x.([]int))
}
func (e *event) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
