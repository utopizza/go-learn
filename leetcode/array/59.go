package array

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	left, right, up, down := 0, n-1, 0, n-1
	m := 1
	for {
		for i := left; i <= right; i++ {
			matrix[up][i] = m
			m++
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			matrix[i][right] = m
			m++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			matrix[down][i] = m
			m++
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			matrix[i][left] = m
			m++
		}
		left++
		if left > right {
			break
		}
	}
	return matrix
}
