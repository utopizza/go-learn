package array

func setZeroes(matrix [][]int) {
	rowSet, colSet := make(map[int]bool), make(map[int]bool) // 在第一遍扫描时记录哪些行列有0，第二次扫描时整行/列置0
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowSet[i], colSet[j] = true, true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowSet[i] || colSet[j] {
				matrix[i][j] = 0
			}
		}
	}
}
