package kthlargestelementinanarray

func findKthLargest(nums []int, k int) int {
	return quickSelect(&nums, 0, len(nums)-1, k)
}

func quickSelect(nums *[]int, left, right, k int) int {
	if k > 0 && k <= right-left+1 {
		index := partition(nums, left, right)

		if index-left == k-1 {
			return (*nums)[index]
		}

		if index-left > k-1 {
			return quickSelect(nums, left, index-1, k)
		}

		return quickSelect(nums, index+1, right, k-index+left-1)
	}
	return -1
}

func partition(nums *[]int, left, right int) int {
	num := (*nums)[right]
	i := left
	for j := left; j < right; j++ {
		if (*nums)[j] <= num {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i += 1
		}
	}
	(*nums)[i], (*nums)[right] = (*nums)[right], (*nums)[i]
	return i
}
