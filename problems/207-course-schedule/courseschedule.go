package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesSet := transformEdgesToSet(prerequisites)
	for i := 0; i < numCourses; i++ {
		if !dfs(i, prerequisitesSet) {
			return false
		}
	}
	return true
}

func dfs(v int, edges map[int][]int) bool {
	return dfsUtil(v, edges, map[int]bool{})
}

func dfsUtil(v int, edges map[int][]int, visited map[int]bool) bool {
	result := true
	visited[v] = true
	nextVertices := edges[v]
	if len(nextVertices) != 0 {
		for _, nextV := range nextVertices {
			onStack, isVisted := visited[nextV]

			if !isVisted {
				result = result && dfsUtil(nextV, edges, visited)
			} else if onStack {
				result = result && false
			}
		}
	}
	visited[v] = false
	return result
}

func transformEdgesToSet(edges [][]int) map[int][]int {
	set := map[int][]int{}
	for _, edge := range edges {
		adjs, ok := set[edge[1]]
		if ok {
			set[edge[1]] = append(adjs, edge[0])
		} else {
			set[edge[1]] = []int{edge[0]}
		}
	}
	return set
}
