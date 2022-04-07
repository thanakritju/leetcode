package coinchange

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i-c < 0 {
				break
			}

			if dp[i-c] != math.MaxInt32 {
				dp[i] = min(dp[i], 1+dp[i-c])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
