package array

func spiralOrder(matrix [][]int) []int {
	var solution []int
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	spiralOrderRecursive(matrix, visited, len(matrix), len(matrix[0]), 0, 0, "right", &solution)
	return solution
}

func spiralOrderRecursive(matrix [][]int, visited [][]bool, row, col, i, j int, curDir string, solution *[]int) {
	if i < 0 || i >= row || j < 0 || j >= col || visited[i][j] {
		return
	}
	*solution = append(*solution, matrix[i][j])
	visited[i][j] = true

	var nextDir string
	var nextI, nextJ int
	switch curDir {
	case "right":
		if j+1 < col && !visited[i][j+1] {
			nextDir = curDir
			nextI, nextJ = i, j+1
		} else {
			nextDir = "down"
			nextI, nextJ = i+1, j
		}
	case "down":
		if i+1 < row && !visited[i+1][j] {
			nextDir = curDir
			nextI, nextJ = i+1, j
		} else {
			nextDir = "left"
			nextI, nextJ = i, j-1
		}
	case "left":
		if j-1 >= 0 && !visited[i][j-1] {
			nextDir = curDir
			nextI, nextJ = i, j-1
		} else {
			nextDir = "up"
			nextI, nextJ = i-1, j
		}
	case "up":
		if i-1 >= 0 && !visited[i-1][j] {
			nextDir = curDir
			nextI, nextJ = i-1, j
		} else {
			nextDir = "right"
			nextI, nextJ = i, j+1
		}
	}
	spiralOrderRecursive(matrix, visited, row, col, nextI, nextJ, nextDir, solution)
}
