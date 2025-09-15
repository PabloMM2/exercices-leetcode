func mergeAlternately(word1 string, word2 string) string {
	longer := len(word1)

	if longer < len(word2) {
		longer = len(word2)
	}
	merged := ""
	for i := 0; i < longer; i++ {
		if len(word1) > i {
			merged += string(word1[i])
		}

		if len(word2) > i {
			merged += string(word2[i])
		}
	}
	return merged
}