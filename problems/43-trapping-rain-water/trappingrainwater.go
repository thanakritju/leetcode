package trappingrainwater

func trap(height []int) int {
	count := 0
	foundStartEdge := false
	startEdgeIndex := 0
	endEdgeIndex := 0

	index := 0
	for true {
		if !foundStartEdge {
			if isLocalMaxima(index, height) {
				startEdgeIndex = index
				foundStartEdge = true
			} else {
				index++
			}
		} else {
			endEdgeIndex = findEndEdge(startEdgeIndex, height)
			if endEdgeIndex == 0 {
				break
			}
			count += getWaterBetweenIndexes(startEdgeIndex, endEdgeIndex, height)
			startEdgeIndex = endEdgeIndex
		}
	}

	return count
}

func findEndEdge(startIndex int, height []int) int {
	foundLocalMinima := false
	maxHeightIndex := 0
	maxHeight := 0
	for i := startIndex; i < len(height); i++ {
		if !foundLocalMinima {
			if isLocalMinima(i, height) {
				foundLocalMinima = true
			}
		} else {
			if isLocalMaxima(i, height) {
				if height[i] >= height[startIndex] {
					return i
				} else if height[i] >= maxHeight {
					maxHeightIndex = i
					maxHeight = height[i]
				}
			}
		}
	}
	return maxHeightIndex
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
	for !foundLeft && index-bufferLeft >= 0 {
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
	for !foundRight && index+bufferRight < len(height) {
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
	for !foundLeft && index-bufferLeft >= 0 {
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
	for !foundRight && index+bufferRight < len(height) {
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
