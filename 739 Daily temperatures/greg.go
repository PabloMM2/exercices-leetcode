func dailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))
	stack := [][]int{}

	for key, tem := range temperatures {
		s := []int{tem, key}
		if key == 0 {
			stack = append(stack, s)
			continue
		}

		for i := (len(stack) - 1); i >= 0; i-- {
			if tem > stack[i][0] {
				k := stack[i][1]
				response[k] = key - k
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, s)
	}

	return response
}