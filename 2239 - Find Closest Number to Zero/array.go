func findClosestNumber(nums []int) int {
	pos := 0

	for key, value := range nums {
		aux := value
		min := nums[pos]

		if value < 0 {
			aux *= -1
		}

		if min < 0 {
			min *= -1
		}

		if aux == min && value > nums[pos] {
			pos = key
		} else if aux < min {
			pos = key
		}
	}

	return nums[pos]
}