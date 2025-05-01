package array

func getRow(rowIndex int) []int {
	ret := make([][]int, rowIndex)
	ret[0] = []int{1}
	for i := 1; i < len(ret); i++ {
		ret[i] = make([]int, i+1)
		ret[i][0], ret[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret[rowIndex-1]
}
