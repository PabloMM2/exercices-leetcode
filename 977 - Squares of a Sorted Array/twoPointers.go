func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	res := []int{}

	for l <= r {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			res = append(res, nums[l]*nums[l])
			l++
		} else {
			res = append(res, nums[r]*nums[r])
			r--
		}
	}

	l, r = 0, len(nums)-1
	for l <= r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return res
}