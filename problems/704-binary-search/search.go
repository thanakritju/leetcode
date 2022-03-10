package search

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)/2, len(nums), target)
}

func binarySearch(nums []int, low, middle, high, target int) int {
	if nums[middle] == target {
		return middle
	}
	if low == middle || middle == high {
		return -1
	}
	if nums[middle] > target {
		return binarySearch(nums, low, (low+middle)/2, middle, target)
	} else {
		return binarySearch(nums, middle, (high+middle)/2, high, target)
	}
}
