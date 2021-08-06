package countingbits

func countBits(n int) []int {
	ans := []int{0}
	index := 1
	for index <= n {
		ans = append(ans, getNumberOfOnes(index))
		index++
	}
	return ans
}

func getNumberOfOnes(n int) int {
	count := 0
	for n != 0 {
		count += n % 2
		n = n / 2
	}
	return count
}
