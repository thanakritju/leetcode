package shortestpath

type Node struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	return -1
}

func getNeighbours(node []int) [][]int {
	neighbours := [][]int{}
	for i := range []int{-1, 0, 1} {
		for j := range []int{-1, 0, 1} {
			if i != j {
				neighbours = append(neighbours, []int{node[0] + i, node[1] + j})
			}
		}
	}
	return neighbours
}

func enqueue(queue [][]int, element []int) [][]int {
	return append(queue, element)
}

func dequeue(queue [][]int) ([]int, [][]int) {
	return queue[0], queue[1:]
}
