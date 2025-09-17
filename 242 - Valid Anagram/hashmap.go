func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	s1 := map[rune]int{}
	s2 := map[rune]int{}

	for _, val := range s {
		s1[val] += 1
	}

	for _, val := range t {
		s2[val] += 1
	}

	for key, _ := range s1 {
		if s1[key] != s2[key] {
			return false
		}
	}

	return true
}