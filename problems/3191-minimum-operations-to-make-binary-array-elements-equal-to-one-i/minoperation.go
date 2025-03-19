package minimumoperationstomakebinaryarrayelementsequaltoonei

func minOperations(nums []int) int {
	count := 0

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			nums[i] = (nums[i] + 1) % 2
			nums[i+1] = (nums[i+1] + 1) % 2
			nums[i+2] = (nums[i+2] + 1) % 2
			count += 1
		}
	}

	for _, v := range nums {
		if v == 0 {
			return -1
		}
	}

	return count
}
