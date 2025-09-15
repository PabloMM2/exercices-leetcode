func numJewelsInStones(jewels string, stones string) int {
	ms := map[rune]int{}
	count := 0

	for _, stone := range stones {
		ms[stone] += 1
	}

	for _, jw := range jewels {
		count += ms[jw]
	}

	return count
}