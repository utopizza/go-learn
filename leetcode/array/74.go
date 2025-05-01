package array

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		item := matrix[mid/n][mid%n]
		if target == item {
			return true
		} else if target > item {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
