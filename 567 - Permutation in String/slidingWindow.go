func checkInclusion(s1 string, s2 string) bool {
	var m1 = make([]int, 26)
	var m2 = make([]int, 26)
	l, r := 0, 0

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		m1[rune(s1[i])-97] += 1
		m2[rune(s2[i])-97] += 1
		r++
	}

	if match(m1, m2) {
		return true
	}

	for r < len(s2) {
		m2[rune(s2[l])-97] -= 1
		m2[rune(s2[r])-97] += 1
		l++
		r++

		if match(m1, m2) {
			return true
		}
	}

	return false
}

func match(m1, m2 []int) bool {
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}