func longestCommonPrefix(strs []string) string {

	word := strs[0]
	prefix := ""
	for _, val := range word {
		counter := 0
		prefix += string(val)
		for _, w := range strs {
			reg, _ := regexp.MatchString("^"+prefix+".*", w)

			if reg {
				counter++
			}
		}
		if counter < len(strs) {
			return prefix[:len(prefix)-1]
		}
	}
	return prefix
}