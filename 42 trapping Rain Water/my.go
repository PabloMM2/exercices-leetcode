func trap(height []int) int {

	l, r := 0, 1
	subs := 0
	res := 0

	for r < len(height) {

		if height[r] > height[l] {
			area := ((height[l]) * (r - (l + 1))) - subs
			res += area
			subs = 0
			l = r
			r++
		} else {
			subs += height[r]
			r++
		}
	}

	l = len(height) - 2
	r = len(height) - 1
	subs = 0
	for l >= 0 {
		if height[l] >= height[r] {
			area := ((height[r]) * (r - (l + 1))) - subs
			res += area
			subs = 0
			r = l
			l--
		} else {
			subs += height[l]
			l--
		}
	}
	return res
}