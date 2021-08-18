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
	if c >= 'a' {
		a.array[c-'a'+26]++
	} else {
		a.array[c-'A']++
	}
}

func (a *array) update(s string, oldIndex int, newIndex int, isStartPointer bool) {
	fmt.Printf("Start updating string: %v isStartPointer: %v\n", s, isStartPointer)
	if isStartPointer {
		for i := oldIndex - 1; i >= newIndex; i-- {
			fmt.Printf("Updating at i: %v, oldIndex: %v, newIndex: %v, value: %v\n%v\n", i, oldIndex, newIndex, string(s[i]), a.array)
			a.addChar(s[i])
		}
	} else {
		for i := oldIndex + 1; i <= newIndex; i++ {
			fmt.Printf("Updating at i: %v, oldIndex: %v, newIndex: %v, value: %v\n%v\n", i, oldIndex, newIndex, string(s[i]), a.array)
			a.addChar(s[i])
		}
	}
}

func (a *array) removeChar(c byte) {
	if c >= 'a' {
		a.array[c-'a'+26]--
	} else {
		a.array[c-'A']--
	}
}

func (a array) matches(t array) bool {
	for i := 0; i < 52; i++ {
		if t.array[i] != 0 && a.array[i] < t.array[i] {
			return false
		}
	}
	return true
}

func (a array) getByChar(c byte) int {
	if c >= 'a' {
		return a.array[c-'a'+26]
	}
	return a.array[c-'A']
}

func minWindow(s string, t string) string {
	arrayT := getArray(t)
	arrayS := getArray(s)
	stoppers := getStoppers(s, arrayT)
	fmt.Printf("%v %v\n%v %v\n\n", s, arrayS, t, arrayT)
	if len(stoppers) == 0 {
		return ""
	}

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
		if arrayS.matches(arrayT) {
			return t
		}
		return ""
	}

	for len(stoppers) != 0 || end+1-start >= len(t) {
		fmt.Printf("%v %v %v\n\n", start, end, stoppers)
		subsetLen := end - start
		if startPointerMove {
			if arrayS.matches(arrayT) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
				}
				arrayS.removeChar(s[start])
				start++
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
			if arrayS.matches(arrayT) {
				if subsetLen < minLen {
					ansStart = start
					ansEnd = end
					minLen = subsetLen
				}
				arrayS.removeChar(s[end])
				end--
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
