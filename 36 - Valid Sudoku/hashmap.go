func isValidSudoku(board [][]byte) bool {
	columns := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			value := board[i][j]

			if value == '.' {
				continue
			}

			if rows[i][value] {
				return false
			}
			rows[i][value] = true

			if columns[j][value] {
				return false
			}
			columns[j][value] = true

			boxId := (i/3)*3 + (j / 3)

			if box[boxId][value] {
				return false
			}
			box[boxId][value] = true
		}
	}

	return true
}