package converttotitle

func convertToTitle(columnNumber int) string {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ans := ""
	for columnNumber != 0 {
		remainder := columnNumber % 26
		columnNumber = (columnNumber - 1) / 26

		if remainder-1 < 0 {
			remainder += 26
		}
		ans = string(s[remainder-1]) + ans
	}
	return ans
}
