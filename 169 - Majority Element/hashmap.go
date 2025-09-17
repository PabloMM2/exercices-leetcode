func majorityElement(nums []int) int {
	mn := map[int]int{}
	min := 0
	response := 0

	for i := 0; i < len(nums); i++ {
		mn[nums[i]] += 1
	}

	for key, val := range mn {
		if val > min {
			min = val
			response = key
		}
	}

	return response
}