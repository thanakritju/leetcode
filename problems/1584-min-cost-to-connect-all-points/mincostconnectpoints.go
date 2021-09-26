package mincosttoconnectallpoints

import "math"

func minCostConnectPoints(points [][]int) int {
	return 0
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

type graph struct {
	edges map[int][]edge
}

type edge struct {
	nodeID int
	weight int
	next   int
}
