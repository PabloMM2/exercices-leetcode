func groupAnagrams(strs []string) [][]string {

	ms := map[string][]int{}
	response := [][]string{}

	for key, val := range strs {
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		ss := string(runes)
		ms[ss] = append(ms[ss], key)
	}

	for _, values := range ms {
		col := []string{}
		for _, val := range values {
			col = append(col, strs[val])
		}
		response = append(response, col)
	}
	return response
}