func lengthOfLongestSubstring(s string) int {
	char := make(map[byte]int)
	l, r := 0, 0
	max, count := 0, 0
	for r < len(s) {

		if key, ok := char[s[r]]; ok {
			for l <= key {
				delete(char, s[l])
				l++
				count--
			}
			continue
		}

		char[s[r]] = r
		count++
		r++

		if count > max {
			max = count
		}
	}
	return max
}