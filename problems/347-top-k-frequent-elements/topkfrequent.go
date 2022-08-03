package topkfrequentelements

import "container/heap"

type item struct {
	val  int
	size int
}

type pq []item

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].size > p[j].size
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
	*p = append(*p, x.(item))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	res := &pq{}
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			m[num] += 1
		} else {
			m[num] = 1
		}
	}

	for key, value := range m {
		heap.Push(res, item{val: key, size: value})
	}

	answer := []int{}
	for i := 0; i < k; i++ {
		answer = append(answer, heap.Pop(res).(item).val)
	}
	return answer
}
