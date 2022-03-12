package nextpermutation

func nextPermutation(nums []int) {
	n := len(nums)
	k := -1
	for x := 0; x < n-1; x++ {
		if nums[x] < nums[x+1] {
			if x > k {
				k = x
			}
		}
	}
	if k > -1 {
		i := -1
		for x := 0; x < n; x++ {
			if nums[k] < nums[x] {
				if x > i {
					i = x
				}
			}
		}
		swap(nums, k, i)
		reverse(nums[k+1:])
	} else {
		reverse(nums)
	}
}

func reverse(nums []int) {
	n := len(nums)
	for x := 0; x < n/2; x++ {
		swap(nums, x, n-x-1)
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
