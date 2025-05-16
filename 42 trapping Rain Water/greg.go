func trap(height []int) int {
	n := len(height)
	l_wall, r_wall := 0, 0
	max_left, max_right := make([]int, n), make([]int, n)
	j := n - 1

	for i := 0; i < n; i++ {
		max_left[i] = l_wall
		max_right[j] = r_wall

		if height[i] > l_wall {
			l_wall = height[i]
		}

		if height[j] > r_wall {
			r_wall = height[j]
		}
		j--
	}

	sum := 0

	for i := 0; i < n; i++ {
		pot := 0
		if max_left[i] < max_right[i] {
			pot = max_left[i]
		} else {
			pot = max_right[i]
		}

		res := pot - height[i]

		if res > 0 {
			sum += res
		}
	}

	return sum
}