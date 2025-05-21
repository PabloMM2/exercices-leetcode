func firstBadVersion(n int) int {
	l, r := 1, n

	for l <= r {
		if isBadVersion(l) {
			return l
		}
		m := int((l + r) / 2)
		if !isBadVersion(m) {
			l = m + 1
		} else {
			r = m
		}
	}

	return 1
}