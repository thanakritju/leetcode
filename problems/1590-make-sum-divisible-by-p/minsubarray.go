package minsubarray

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	need := sum % p
	cur := 0
	dp := map[int]int{
		0: -1,
	}
	res := len(nums)

	for i, num := range nums {
		cur = (cur + num) % p
		dp[cur] = i
		t := (cur - need) % p
		if t < 0 {
			t += p
		}
		index, ok := dp[t]
		if ok {
			res = min(res, i-index)
		}
	}

	if res < len(nums) {
		return res
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
