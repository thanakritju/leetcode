package minimumwindowsubstring

import "fmt"

type array struct {
	array [52]int
}

func getArray(s string) array {
	var a array
	for i := 0; i < len(s); i++ {
		a.addChar(s[i])
	}
	return a
}

func (a *array) addChar(c byte) {
	a.array[charToIndex(c)]++
}

func (a *array) update(s string, oldIndex int, newIndex int, isStartPointer bool) {
	if isStartPointer {
		for i := oldIndex - 1; i >= newIndex; i-- {
			a.addChar(s[i])
		}
	} else {
		for i := oldIndex + 1; i <= newIndex; i++ {
			a.addChar(s[i])
		}
	}
}

func (a *array) removeChar(c byte) {
	a.array[charToIndex(c)]--
}

func (a array) matches(t array) bool {
	for i := 0; i < 52; i++ {
		if t.array[i] != 0 && a.array[i] < t.array[i] {
			return false
		}
	}
	return true
}

func charToIndex(c byte) int {
	if c >= 'a' {
		return int(c - 'a' + 26)
	}
	return int(c - 'A')
}

func (a array) check(t array, c byte) bool {
	index := charToIndex(c)
	if t.array[index] == 0 {
		return true
	}
	return a.array[index] >= t.array[index]
}

func (a array) getByChar(c byte) int {
	return a.array[charToIndex(c)]
}

func minWindow(s string, t string) string {
	arrayT := getArray(t)
	arrayS := getArray(s)
	stoppers := getStoppers(s, arrayT)
	if len(stoppers) == 0 {
		return ""
	}

	start := 0
	end := len(s) - 1
	start, stoppers = stoppers[0], stoppers[1:]
	end, stoppers = stoppers[len(stoppers)-1], stoppers[:len(stoppers)-1]

	minLen := len(s)
	ansStart := 0
	ansEnd := -1

	startPointerMove := true

	if len(t) > len(s) {
		return ""
	}

	if arrayS.matches(arrayT) {
		minLen = len(s)
		ansStart = 0
		ansEnd = len(s) - 1
	} else {
		return ""
	}

	if len(t) == 1 {
		return t
	}

	for end+1-start > len(t) {
		fmt.Printf("start %v end %v %v\n", start, end, stoppers)
		subsetLen := end - start
		if startPointerMove {
			arrayS.removeChar(s[start])
			start++
			if arrayS.check(arrayT, s[start-1]) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
					fmt.Printf("saving value %v end+1-start %v len(t) %v startPointerMove %v\n", s[ansStart:ansEnd+1], end+1-start, len(t), startPointerMove)
				} else if minLen == subsetLen && ansEnd == end && ansStart == start {
					break
				}
			} else {
				if len(stoppers) == 0 {
					break
				}
				oldStart := start
				start, stoppers = stoppers[0], stoppers[1:] // pop left
				arrayS.update(s, oldStart, start, true)
				startPointerMove = !startPointerMove
			}
		} else {
			arrayS.removeChar(s[end])
			end--
			if arrayS.check(arrayT, s[end+1]) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
					fmt.Printf("saving value %v end+1-start %v len(t) %v startPointerMove %v\n", s[ansStart:ansEnd+1], end+1-start, len(t), startPointerMove)
				} else if minLen == subsetLen && ansEnd == end && ansStart == start {
					break
				}
			} else {
				if len(stoppers) == 0 {
					break
				}
				oldEnd := end
				end, stoppers = stoppers[len(stoppers)-1], stoppers[:len(stoppers)-1] // pop right
				arrayS.update(s, oldEnd, end, false)
				startPointerMove = !startPointerMove
			}
		}
	}

	return s[ansStart : ansEnd+1]
}

// "ZZADOBECODEBANCZZ", "ABC" -> [2 5 7 11 12 14]
func getStoppers(s string, t array) []int {
	arr := []int{}
	for i := 0; i < len(s); i++ {
		if t.getByChar(s[i]) > 0 {
			arr = append(arr, i)
		}
	}
	return arr
}
