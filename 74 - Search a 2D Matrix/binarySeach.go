func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix[i])-1
		for l <= r {
			m := int((l + r) / 2)
			if target > matrix[i][m] {
				l = m + 1
			} else if target < matrix[i][m] {
				r = m - 1
			} else if target == matrix[i][m] {
				return true
			}
		}
	}
	return false
}