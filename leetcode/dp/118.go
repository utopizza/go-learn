package dp

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		var row []int
		row = append(row, 1)
		for j := 1; j < i; j++ {
			sum := result[i-1][j-1] + result[i-1][j]
			row = append(row, sum)
		}
		row = append(row, 1)
		result[i] = row
	}
	return result
}
