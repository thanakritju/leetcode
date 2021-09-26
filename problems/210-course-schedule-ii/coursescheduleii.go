package coursescheduleii

var adjMatrix [2000][2000]int
var sortedList []int
var foundCyclic bool

const (
	NOT_VISITED = iota
	TEMPORARY
	PERMANENT
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjMatrix = [2000][2000]int{}
	for _, v := range prerequisites {
		adjMatrix[v[1]][v[0]] = 1
	}
	visited := make(map[int]int)
	sortedList = []int{}
	foundCyclic = false

	for i := 0; i < numCourses; i++ {
		_, v := visited[i]
		if v {
			continue
		}
		if !foundCyclic {
			visit(i, numCourses, visited)
		}
	}

	if foundCyclic {
		return []int{}
	} else {
		return sortedList
	}
}

func visit(n int, numCourses int, visited map[int]int) {
	v := visited[n]
	if v == PERMANENT {
		return
	}
	if v == TEMPORARY {
		foundCyclic = true
		return
	}
	visited[n] = TEMPORARY

	for i := 0; i < numCourses; i++ {
		if adjMatrix[n][i] == 1 {
			visit(i, numCourses, visited)
		}
	}

	visited[n] = PERMANENT
	sortedList = append([]int{n}, sortedList...)
}
