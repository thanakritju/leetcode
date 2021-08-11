package isgraphbipartite

func isBipartite(graph [][]int) bool {
	visited := map[int]bool{}
	for i := 0; i < len(graph); i++ {
		dfs(i, graph, visited, true)
	}
	for v, adjs := range graph {
		for _, adj := range adjs {
			if visited[v] == visited[adj] {
				return false
			}
		}
	}
	return true
}

func dfs(v int, graph [][]int, visited map[int]bool, shouldGoToSetA bool) {
	_, isVisted := visited[v]
	if !isVisted {
		visited[v] = shouldGoToSetA
	}
	for _, adj := range graph[v] {
		_, isVisted := visited[adj]
		if !isVisted {
			dfs(adj, graph, visited, !shouldGoToSetA)
		}
	}
}
