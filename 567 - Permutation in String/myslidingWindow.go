func checkInclusion(s1 string, s2 string) bool {
	var m1 = make(map[rune]int)
	var m2 = make(map[rune]int)
	l, r := 0, 0

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		m1[rune(s1[i])]++
		m2[rune(s2[i])]++
		r++
	}

	if match(m1, m2) {
		return true
	}

	for r < len(s2) {
		m2[rune(s2[l])]--
		m2[rune(s2[r])]++
		l++
		r++

		if match(m1, m2) {
			return true
		}
	}

	return false
}

func match(m1, m2 map[rune]int) bool {
	for key, val := range m1 {
		v, ok := m2[key]
		if !ok {
			return false
		} else if ok && v != val {
			return false
		}
	}
	return true
}