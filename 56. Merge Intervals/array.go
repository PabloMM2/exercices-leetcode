func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	response := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			response = append(response, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}

	response = append(response, []int{start, end})

	return response
}