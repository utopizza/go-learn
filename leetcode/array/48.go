package array

func rotate(matrix [][]int) {
	row, column := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := i; j < column; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column/2; j++ {
			matrix[i][j], matrix[i][column-1-j] = matrix[i][column-1-j], matrix[i][j]
		}
	}
}
