package hash

func solveSudoku(board [][]byte) {
	// 用这几个数组记录某行、某列、某9格已经用过的数字
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	// 初始化填充
	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				spaces = append(spaces, [2]int{i, j}) // 空白格的坐标记录下来
			} else {
				digit := v - '1'
				// 标记digit已经使用
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	// 回溯填充
	solveSudokuDFS(board, 0, line, column, block, spaces)
}

func solveSudokuDFS(board [][]byte, pos int, line, column [9][9]bool, block [3][3][9]bool, spaces [][2]int) bool {
	if pos == len(spaces) {
		return true
	}
	i, j := spaces[pos][0], spaces[pos][1]
	for digit := byte(0); digit < 9; digit++ {
		if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] { // 若果digit未用过

			// 标记digit已使用
			line[i][digit] = true
			column[j][digit] = true
			block[i/3][j/3][digit] = true

			// 填入digit
			board[i][j] = digit + '1'

			// dfs下一空格
			if solveSudokuDFS(board, pos+1, line, column, block, spaces) {
				return true
			}

			// 回溯，取消标记digit已经使用
			line[i][digit] = false
			column[j][digit] = false
			block[i/3][j/3][digit] = false
		}
	}
	return false
}
