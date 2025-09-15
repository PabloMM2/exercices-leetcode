func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	if r == 0 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	for l <= r {

		m := ((l + r) / 2)

		if nums[m] == target {
			return m
		} else if nums[m] > target {

			if nums[l] <= target || nums[m] < nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}

		} else if nums[m] < target {

			if nums[r] >= target || nums[m] > nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}

	return -1
}
