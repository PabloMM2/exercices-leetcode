func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	min := len(nums) + 1
	count := 0

	for r < len(nums) {
		count += nums[r]
		for count >= target {
			size := (r - l) + 1
			if size < min {
				min = size
			}
			count -= nums[l]
			l++
		}

		r++
	}

	if min > len(nums) {
		return 0
	}

	return min
}