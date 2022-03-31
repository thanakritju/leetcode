package perfectsquares

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for target := 1; target*target <= n; target++ {
		for s := target * target; s <= n; s++ {
			dp[s] = min(dp[s], 1+dp[s-target*target])
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
