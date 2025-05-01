package array

func isValidSudoku(board [][]byte) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	// 行检测
	for i := 0; i < len(board); i++ {
		m := make(map[byte]bool)
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if m[board[i][j]] {
				return false
			}
			m[board[i][j]] = true
		}
	}
	// 列检测
	for j := 0; j < len(board[0]); j++ {
		m := make(map[byte]bool)
		for i := 0; i < len(board); i++ {
			if board[i][j] == '.' {
				continue
			}
			if m[board[i][j]] {
				return false
			}
			m[board[i][j]] = true
		}
	}
	// 9格检测
	var maps [3][3]map[byte]bool
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			maps[i][j] = make(map[byte]bool)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if maps[i/3][j/3][board[i][j]] {
				return false
			}
			maps[i/3][j/3][board[i][j]] = true
		}
	}
	return true
}
