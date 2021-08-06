package trappingrainwater

func trap(height []int) int {
	return 0
}

func getWaterBetweenIndexes(index1 int, index2 int, height []int) int {
	var maxHeight int
	if height[index1] > height[index2] {
		maxHeight = height[index2]
	} else {
		maxHeight = height[index1]
	}
	count := 0
	for i := index1; i < index2; i++ {
		count += maxHeight - height[i]
	}

	return count
}
