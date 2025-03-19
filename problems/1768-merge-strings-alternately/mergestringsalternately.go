package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {
	ans := ""

	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			ans += string(word1[i])
		}
		if i < len(word2) {
			ans += string(word2[i])
		}
		i++
	}

	return ans
}
