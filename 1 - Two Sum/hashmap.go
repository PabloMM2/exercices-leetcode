func twoSum(nums []int, target int) []int {

	mn := map[int]int{}
	response := []int{}

	for key, val := range nums {
		mn[val] = key
	}

	for key, val := range nums {
		pos, ok := mn[target-val]
		if ok && pos != key {
			response = append(response, key)
			response = append(response, pos)
			return response
		}
	}

	return response
}