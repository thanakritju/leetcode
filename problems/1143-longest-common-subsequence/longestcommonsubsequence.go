package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := [1001][1001]int{}
	l1 := len(text1)
	l2 := len(text2)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[l1][l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
