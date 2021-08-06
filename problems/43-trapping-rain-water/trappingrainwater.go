package trappingrainwater

func trap(height []int) int {
	count := 0
	return count
}

func isLocalMaxima(index int, height []int) bool {
	if index == 0 {
		return height[index] > height[index+1]
	}
	if index == len(height)-1 {
		return height[index-1] < height[index]
	}
	return height[index-1] < height[index] && height[index] > height[index+1]
}

func isLocalMinima(index int, height []int) bool {
	if index == 0 {
		return height[index] < height[index+1]
	}
	if index == len(height)-1 {
		return height[index-1] > height[index]
	}
	return height[index-1] > height[index] && height[index] < height[index+1]
}

func getWaterBetweenIndexes(index1 int, index2 int, height []int) int {
	println(index1, index2)
	var maxHeight int
	if height[index1] > height[index2] {
		maxHeight = height[index2]
	} else {
		maxHeight = height[index1]
	}
	count := 0
	for i := index1; i < index2; i++ {
		if maxHeight >= height[i] {
			count += maxHeight - height[i]
		}
	}

	return count
}
