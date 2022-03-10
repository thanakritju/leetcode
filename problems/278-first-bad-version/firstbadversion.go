package firstbadversion

func firstBadVersion(n int) int {
	l := 0
	r := n - 1
	for l < r {
		m := (l + r) / 2
		if isBadVersion(m) {
			if !isBadVersion(m - 1) {
				return m + 1
			} else {
				r = m - 1
			}
		} else {
			l = m + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	return []bool{false, false, false, true, true}[version]
}
