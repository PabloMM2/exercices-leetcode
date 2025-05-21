func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	max, count, replace := 0, 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			if replace < k {
				replace++
			} else {
				for replace >= k {
					if nums[l] == 0 {
						replace--
					}
					count--
					l++
				}
				continue
			}
		}
		count++
		r++

		if count > max {
			max = count
		}
	}

	return max
}