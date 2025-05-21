func isPerfectSquare(num int) bool {
	l, r := 0, num

	for l <= r {
		m := int((l + r) / 2)
		sqr := m * m

		if sqr == num {
			return true
		} else if sqr > num {
			r = m - 1
		} else if sqr < num {
			l = m + 1
		}
	}
	return false
}