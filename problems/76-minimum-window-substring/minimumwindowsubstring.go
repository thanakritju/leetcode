package minimumwindowsubstring

func minWindow(s string, t string) string {
	asciiTable := getTable(t)
	stoppers := getStoppers(s, asciiTable)

	start := 0
	end := len(s) - 1

	minLen := len(s)
	ansStart := 0
	ansEnd := -1

	startPointerMove := true

	if len(t) > len(s) {
		return ""
	}

	if len(t) == 1 {
		if areSubset(s, t) {
			return t
		}
		return ""
	}

	for len(stoppers) != 0 || end+1-start >= len(t) {
		subsetLen := end - start
		currentString := s[start : end+1]
		if startPointerMove {
			if areSubset(currentString, t) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
				}
				start++
			} else {
				if len(stoppers) == 0 {
					break
				}
				start, stoppers = stoppers[0], stoppers[1:] // pop left
				startPointerMove = !startPointerMove
			}
		} else {
			if areSubset(currentString, t) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
				}
				end--
			} else {
				if len(stoppers) == 0 {
					break
				}
				end, stoppers = stoppers[len(stoppers)-1], stoppers[:len(stoppers)-1] // pop right
				startPointerMove = !startPointerMove
			}
		}
	}

	return s[ansStart : ansEnd+1]
}

// "ZZADOBECODEBANCZZ", "ABC" -> [2 5 7 11 12 14]
func getStoppers(s string, asciiTable [128]int) []int {
	arr := []int{}
	for index, char := range s {
		if asciiTable[int(char)] < 0 {
			arr = append(arr, index)
		}
	}
	return arr
}

func areSubset(s string, t string) bool {
	asciiTable := getTable(t)
	for i := 0; i < len(s); i++ {
		asciiTable[int(s[i])]++
	}
	for _, value := range asciiTable {
		if value < 0 {
			return false
		}
	}
	return true
}

func getTable(t string) [128]int {
	var asciiTable [128]int
	for i := 0; i < len(t); i++ {
		asciiTable[int(t[i])]--
	}
	return asciiTable
}
