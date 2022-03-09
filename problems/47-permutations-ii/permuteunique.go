package permuteunique

func permuteUnique(nums []int) [][]int {
	return heapGenerate(len(nums), nums, [][]int{})
}

func heapGenerate(k int, nums []int, ans [][]int) [][]int {
	if k == 1 {
		ans = merge(ans, nums)
	} else {
		ans = heapGenerate(k-1, nums, ans)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				temp := nums[k-1]
				nums[k-1] = nums[i]
				nums[i] = temp
			} else {
				temp := nums[k-1]
				nums[k-1] = nums[0]
				nums[0] = temp
			}
			ans = heapGenerate(k-1, nums, ans)
		}
	}

	return ans
}

func merge(ans [][]int, current []int) [][]int {
	for _, each := range ans {
		if equal(each, current) {
			return ans
		}
	}
	tmp := make([]int, len(current))
	copy(tmp, current)
	return append(ans, tmp)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
