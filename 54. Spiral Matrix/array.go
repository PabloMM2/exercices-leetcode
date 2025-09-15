func spiralOrder(matrix [][]int) []int {
	height := len(matrix)
	widht := len(matrix[0])
	direction := "rigth"
	response := []int{}
	ll := 0
	lr := widht - 1
	lu := 0
	ld := height - 1

	for len(response) < (height * widht) {
		switch direction {
		case "rigth":
			for i := ll; i <= lr; i++ {
				response = append(response, matrix[ll][i])
			}
			lu += 1
			direction = "down"
		case "down":
			for i := lu; i <= ld; i++ {
				response = append(response, matrix[i][lr])
			}
			lr -= 1
			direction = "left"
		case "left":
			for i := lr; i >= ll; i-- {
				response = append(response, matrix[ld][i])
			}
			ld -= 1
			direction = "up"
		case "up":
			for i := ld; i >= lu; i-- {
				response = append(response, matrix[i][ll])
			}
			ll += 1
			direction = "rigth"
		}
	}

	return response
}