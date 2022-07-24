package searchmatrix

func searchMatrix(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		if binarySearch(arr, target) {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	min := 0
	max := len(arr)

	for min < max {
		mid := (min + max) / 2
		if target == arr[mid] {
			return true
		}
		if target > arr[mid] {
			min = mid + 1
		}
		if target < arr[mid] {
			max = mid
		}
	}

	return false
}
