func summaryRanges(nums []int) []string {

	response := []string{}

	if len(nums) == 0 {
		return response
	}

	values := fmt.Sprintf("%d", nums[0])
	last := nums[0]

	if len(nums) == 1 {
		response = append(response, values)
	}

	for i := 1; i < len(nums); i++ {

		if (last + 1) == nums[i] {
			last = nums[i]
		} else {

			if values != fmt.Sprintf("%d", last) {
				values += fmt.Sprintf("->%d", last)
			}

			response = append(response, values)
			last = nums[i]
			values = fmt.Sprintf("%d", last)

		}

		if i == len(nums)-1 {

			if values != fmt.Sprintf("%d", last) {
				values += fmt.Sprintf("->%d", last)
			}
			response = append(response, values)
		}

	}
	return response
}