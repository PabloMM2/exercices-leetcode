func lengthOfLongestSubstring(s string) int {
	mapOfChar := map[rune]int{}
	var maxValue int
	var lastPosition int

	for i, value := range s {

		key, isInMap := mapOfChar[value]
		if isInMap {
			if maxValue < (i - lastPosition) {
				maxValue = (i - lastPosition)
			}

			if (key + 1) > lastPosition {
				lastPosition = key + 1
			}
		}

		if i+1 == len(s) {
			if maxValue < (len(s) - lastPosition) {
				maxValue = (len(s) - lastPosition)
			}
		}

		mapOfChar[value] = i
	}

	return maxValue
}