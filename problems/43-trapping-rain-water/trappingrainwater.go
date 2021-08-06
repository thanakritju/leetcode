package trappingrainwater

func trap(height []int) int {
	count := 0
	return count
}

func isLocalMaxima(index int, height []int) bool {
	foundLeft := false
	if index == 0 {
		foundLeft = true
	}
	foundRight := false
	if index == len(height)-1 {
		foundRight = true
	}
	bufferLeft := 1
	for !foundLeft {
		if height[index-bufferLeft] == height[index] {
			bufferLeft++
		} else {
			foundLeft = true
			if height[index-bufferLeft] > height[index] {
				return false
			}
		}
	}
	bufferRight := 1
	for !foundRight {
		if height[index] == height[index+bufferRight] {
			bufferRight++
		} else {
			foundRight = true
			if height[index] < height[index+bufferRight] {
				return false
			}
		}
	}
	return true
}

func isLocalMinima(index int, height []int) bool {
	foundLeft := false
	if index == 0 {
		foundLeft = true
	}
	foundRight := false
	if index == len(height)-1 {
		foundRight = true
	}
	bufferLeft := 1
	for !foundLeft {
		if height[index-bufferLeft] == height[index] {
			bufferLeft++
		} else {
			foundLeft = true
			if height[index-bufferLeft] < height[index] {
				return false
			}
		}
	}
	bufferRight := 1
	for !foundRight {
		if height[index] == height[index+bufferRight] {
			bufferRight++
		} else {
			foundRight = true
			if height[index] > height[index+bufferRight] {
				return false
			}
		}
	}
	return true
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
