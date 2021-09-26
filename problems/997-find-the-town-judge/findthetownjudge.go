package findthetownjudge

func findJudge(n int, trust [][]int) int {
	var adjMatrix [1001][1001]int

	for _, v := range trust {
		adjMatrix[v[0]][v[1]] = 1
	}

	for i := 1; i <= n; i++ {
		sumX := 0
		for _, v := range adjMatrix[i] {
			sumX += v
		}
		if sumX == 0 {
			sumY := 0
			for j := 1; j <= n; j++ {
				sumY += adjMatrix[j][i]
			}
			if sumY == n-1 {
				return i
			}
		}
	}
	return -1
}
