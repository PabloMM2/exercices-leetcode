func longestConsecutive(nums []int) int {

	var mapN = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mapN[nums[i]] = true
	}

	max := 0
	for key, _ := range mapN {
		if !mapN[key-1] {
			current := key
			count := 1
			for mapN[current+1] {
				current++
				count++
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}