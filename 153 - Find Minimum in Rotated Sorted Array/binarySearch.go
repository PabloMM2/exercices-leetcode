func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	if r == 0 {
		return nums[0]
	}

	for l < r {
		m := int((l + r) / 2)

		if nums[m] >= nums[l] {
			if nums[m+1] < nums[m] {
				return nums[m+1]
			}
			l = m + 1
		} else if nums[l] > nums[m] {
			r = m
		}
	}
	return nums[0]
}