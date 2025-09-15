func canConstruct(ransomNote string, magazine string) bool {
	mg := map[rune]int{}
	rn := map[rune]int{}

	for _, note := range ransomNote {
		rn[note] += 1
	}

	for _, letter := range magazine {
		mg[letter] += 1
	}

	for key, value := range rn {
		if mg[key] < value {
			return false
		}
	}
	return true
}