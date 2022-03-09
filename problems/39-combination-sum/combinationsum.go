package combinationsum

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Printf("candidates: %v", candidates)
	return findSolutions([]int{}, candidates, target, [][]int{})
}

func findSolutions(current, candidates []int, target int, ans [][]int) [][]int {
	for _, each := range candidates {
		sum := sum(current) + each
		if sum < target {
			ans = findSolutions(append(current, each), candidates, target, ans)
		}
		if sum == target {
			ans = merge(ans, append(current, each))
		}
	}
	return ans
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func merge(ans [][]int, current []int) [][]int {
	sort.Ints(current)
	for _, each := range ans {
		if arePermutation(each, current) {
			return ans
		}
	}
	return append(ans, current)
}

func arePermutation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	var table [26]int
	for i := 0; i < len(a); i++ {
		table[a[i]]++
		table[b[i]]--
	}
	for _, value := range table {
		if value != 0 {
			return false
		}
	}
	return true
}
