package combinations

func combine(n int, k int) [][]int {
	ans := [][]int{}
	for i := 1; i <= n; i++ {
		ans = combineRecur([]int{i}, ans, n, k)
	}
	return ans
}

func combineRecur(arr []int, answer [][]int, n, k int) [][]int {
	if len(arr) == k {
		answer = append(answer, arr)
	} else {
		for i := arr[len(arr)-1] + 1; i <= n; i++ {
			newArr := getCopy(arr)
			answer = combineRecur(append(newArr, i), answer, n, k)
		}
	}
	return answer
}

func getCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
