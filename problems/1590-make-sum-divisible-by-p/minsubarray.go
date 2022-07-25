package minsubarray

func minSubarray(nums []int, p int) int {
	count = 0
	if divide(nums, p) {

	}
	return 0
}

func search(nums []int, p int, count int) bool {

}

func divide(nums []int, p int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum%p == 0
}
