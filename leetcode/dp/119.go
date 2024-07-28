package dp

func getRow(rowIndex int) []int {
	row := []int{1}
	for i := 1; i <= rowIndex; i++ {
		var nextRow []int
		nextRow = append(nextRow, 1)
		for j := 1; j < len(row); j++ {
			nextRow = append(nextRow, row[j-1]+row[j])
		}
		nextRow = append(nextRow, 1)
		row = nextRow
	}
	return row
}
