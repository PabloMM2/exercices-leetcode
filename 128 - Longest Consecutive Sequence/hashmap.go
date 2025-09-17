func longestConsecutive(nums []int) int {
	mn := map[int]bool{}
	response := 0

	for i := 0; i < len(nums); i++ {
		mn[nums[i]] = true
	}

	for key, _ := range mn {
		count := 1
		if _, ok := mn[key-1]; !ok {
			for {
				key += 1
				if _, ok := mn[key]; !ok {
					break
				}
				count += 1
			}
		}

		if count > response {
			response = count
		}
	}

	return response
}