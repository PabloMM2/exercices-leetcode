func longestConsecutive(nums []int) int {

	var mapN = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mapN[nums[i]]++
	}

	var sortNums []int
	for key, _ := range mapN {
		sortNums = append(sortNums, key)
	}
	sort.Ints(sortNums)

	count := 0
	last := 0
	max := 0
	for _, val := range sortNums {
		if count == 0 {
			last = val
			count++
			if max < count {
				max = count
			}
			continue
		}

		if (last + 1) == val {
			count++
			if max < count {
				max = count
			}
		} else {
			if max < count {
				max = count
			}
			count = 1
		}

		last = val
	}

	return max
}