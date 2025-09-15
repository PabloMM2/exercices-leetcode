func containsDuplicate(nums []int) bool {
	d := map[int]bool{}
	for _, num := range nums {
		if d[num] {
			return true
		}
		d[num] = true
	}
	return false
}