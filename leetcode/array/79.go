package array

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existBackTrack(board, word, 0, i, j, visited) {
				return true
			}
		}
	}
	return false
}

func existBackTrack(board [][]byte, word string, charIndex int, i, j int, visited [][]bool) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[charIndex] || visited[i][j] {
		return false
	}
	if charIndex == len(word)-1 {
		return true
	}
	visited[i][j] = true
	found := existBackTrack(board, word, charIndex+1, i-1, j, visited) ||
		existBackTrack(board, word, charIndex+1, i+1, j, visited) ||
		existBackTrack(board, word, charIndex+1, i, j-1, visited) ||
		existBackTrack(board, word, charIndex+1, i, j+1, visited)
	visited[i][j] = false
	return found
}
