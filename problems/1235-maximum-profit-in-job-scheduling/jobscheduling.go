package maximumprofitinjobscheduling

import "sort"

type Job struct {
	startTime int
	endTime   int
	profit    int
}
type Jobs []Job

func (j Jobs) Len() int { return len(j) }
func (j Jobs) Less(x, y int) bool {
	return j[x].startTime < j[y].startTime
}
func (j Jobs) Swap(x, y int) { j[x], j[y] = j[y], j[x] }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var jobs Jobs
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, Job{startTime: startTime[i], endTime: endTime[i], profit: profit[i]})
	}
	sort.Sort(jobs)
	return len(jobs)
}
