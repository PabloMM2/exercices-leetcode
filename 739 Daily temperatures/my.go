func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	max, position, last := 0, 0, 0
	for i := 0; i < len(temperatures); i++ {

		if temperatures[i] < max && position > i && last <= temperatures[i] {
			stack = append(stack, position-i)
			last = temperatures[i]
			continue
		}

		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				max = temperatures[j]
				position = j
				stack = append(stack, j-i)
				break
			}
			if j+1 == len(temperatures) {
				stack = append(stack, 0)
			}
		}

		if i+1 == len(temperatures) {
			stack = append(stack, 0)
			continue
		}

		last = temperatures[i]
	}
	return stack
}