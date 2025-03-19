package greatestcommondivisorofstrings

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	l := gcd(len(str1), len(str2))

	return str1[0:l]
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
