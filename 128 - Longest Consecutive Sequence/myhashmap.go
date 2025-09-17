func longestConsecutive(nums []int) int {

	mn := map[int]bool{}
	response := 0

	for i := 0; i < len(nums); i++ {
		mn[nums[i]] = true
	}

	for key, val := range mn {
		if !val {
			continue
		}
		count := 1
		mn[key] = false
		for {
			key += 1
			if _, ok := mn[key]; !ok {
				break
			}
			count += 1
			mn[key] = false
		}

		if count > response {
			response = count
		}
	}

	return response
}