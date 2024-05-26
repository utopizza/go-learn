package leetcode

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// 行列转置
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 左右翻转
	col := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < col/2; j++ {
			matrix[i][j], matrix[i][col-1-j] = matrix[i][col-1-j], matrix[i][j]
		}
	}
}
