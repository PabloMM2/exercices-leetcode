func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	word := strs[0]
	response := ""
	for key, val := range word {

		for _, w := range strs {
			if string(w[key]) != string(val) {
				return response
			}
		}
		response += string(val)
	}
	return response
}