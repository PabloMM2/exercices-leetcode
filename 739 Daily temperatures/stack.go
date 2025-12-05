func dailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))
	stack := []int{}

	for key, val := range temperatures {

		for len(stack) > 0 && val > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			response[last] = key - last
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, key)
	}

	return response
}