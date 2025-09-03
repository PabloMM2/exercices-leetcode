func romanToInt(s string) int {
	rn := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	total := rn[string(s[len(s)-1])]
	last := string(s[len(s)-1])
	for i := len(s) - 2; i >= 0; i-- {

		if rn[string(s[i])] >= rn[last] {
			total += rn[string(s[i])]
		} else {
			total -= rn[string(s[i])]
		}

		last = string(s[i])
	}

	return total
}