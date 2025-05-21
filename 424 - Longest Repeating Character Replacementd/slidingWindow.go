func characterReplacement(s string, k int) int {
	l, r := 0, 0
	long := 0
	count := make([]int, 26)

	for r < len(s) {
		count[rune(s[r])-65] += 1
		for ((r-l)+1)-Max(count) > k {
			count[rune(s[l])-65] -= 1
			l++
		}

		if (r-l)+1 > long {
			long = (r - l) + 1
		}
		r++
	}
	return long
}

func Max(count []int) int {
	max := 0
	for i := 0; i < len(count); i++ {
		if count[i] > max {
			max = count[i]
		}
	}
	return max
}