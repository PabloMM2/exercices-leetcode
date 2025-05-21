func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	max, count, replace := 0, 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			if replace < k {
				replace += 1
			} else {
				l += 1
				r = l
				replace, count = 0, 0
				continue
			}
		}
		count += 1
		r += 1

		if count > max {
			max = count
		}
	}

	return max
}