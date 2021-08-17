package findanagrams

func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	arr := []int{}
	for i := 0; i <= lenS-lenP; i++ {
		if arePermutation(s[i:i+lenP], p, lenP) {
			arr = append(arr, i)
		}
	}

	return arr
}

func arePermutation(s1 string, s2 string, lenS int) bool {
	var a [26]int
	for i := 0; i < lenS; i++ {
		a[int(s1[i])-'a']++
		a[int(s2[i])-'a']--
	}
	for _, value := range a {
		if value != 0 {
			return false
		}
	}
	return true
}
