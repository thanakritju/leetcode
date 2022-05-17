package deleteandearn

func deleteAndEarn(nums []int) int {
	var hash = make(map[int]int)
	var memo [100000]int
	maxValue := -1
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
		_, ok := hash[num]
		if ok {
			hash[num] += 1
		} else {
			hash[num] = 1
		}
	}
	for i := 1; i <= maxValue; i++ {
		dp(&memo, hash, i)
	}

	return memo[maxValue]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dp(memo *[100000]int, hash map[int]int, n int) int {
	if n == 1 {
		memo[n] = n * hash[n]
	} else {
		memo[n] = max(n*hash[n]+memo[n-2], memo[n-1])
	}

	return memo[n]
}
