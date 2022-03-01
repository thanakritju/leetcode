package findsubstring

func findSubstring(s string, words []string) []int {
	ans := []int{}
	word_size := len(words[0])
	word_count := len(words)
	size_l := word_size * word_count

	hash := make(map[string]int)
	for _, v := range words {
		_, ok := hash[v]
		if ok {
			hash[v] = hash[v] + 1
		} else {
			hash[v] = 1
		}
	}

	for i := 0; i < len(s) - size_l + 1; i++ {
		temp_hash := make(map[string]int)
		for k, v := range hash {
			temp_hash[k] = v
		}
		j := i
		count := word_count
		for j < i + size_l {
			word := s[j:j+word_size]

			_, ok := hash[word]
			if !ok || temp_hash[word] == 0 {
				break
			} else {
				temp_hash[word] = temp_hash[word] -1
				count -= 1
			}

			j += word_size
		}
		if count == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
