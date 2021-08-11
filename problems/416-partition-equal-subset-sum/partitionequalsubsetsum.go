package partitionequalsubsetsum

func canPartition(nums []int) bool {
	sum := sum(nums)
	if sum%2 != 0 {
		return false
	}
	targetSum := sum / 2
	dp := map[DpKey]bool{}
	return search(dp, nums, 0, targetSum)
}

type DpKey struct {
	targetSum int
	depth     int
}

func search(dp map[DpKey]bool, nums []int, depth int, targetSum int) bool {
	if depth >= len(nums) {
		return 0 == targetSum
	}
	key := DpKey{targetSum: targetSum, depth: depth}
	i, ok := dp[key]
	if ok {
		return i
	}
	result1 := false
	if nums[depth] <= targetSum {
		result1 = search(dp, nums, depth+1, targetSum-nums[depth])
	}
	result2 := search(dp, nums, depth+1, targetSum)
	dp[key] = result1 || result2
	return dp[key]
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
