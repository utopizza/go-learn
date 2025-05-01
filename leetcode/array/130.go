package array

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		solveDFS(board, i, 0)
		solveDFS(board, i, n-1)
	}
	for j := 0; j < n; j++ {
		solveDFS(board, 0, j)
		solveDFS(board, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(board [][]byte, i int, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	solveDFS(board, i+1, j)
	solveDFS(board, i-1, j)
	solveDFS(board, i, j+1)
	solveDFS(board, i, j-1)
}
