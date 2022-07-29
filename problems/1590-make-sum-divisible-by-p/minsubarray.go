package minsubarray

func minSubarray(nums []int, p int) int {
	return 0
}

func divide(nums []int, p int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum%p == 0
}
