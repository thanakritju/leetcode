package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	lenS1 := len(s1)
	lenS2 := len(s2)
	for i := 0; i <= lenS2-lenS1; i++ {
		if arePermutation(s2[i:i+lenS1], s1, lenS1) {
			return true
		}
	}

	return false
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
