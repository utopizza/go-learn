package array

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var found bool
			numIslandsDFS(grid, visited, m, n, i, j, &found)
			if found {
				count++
			}
		}
	}
	return count
}

func numIslandsDFS(grid [][]byte, visited [][]bool, m, n, i, j int, found *bool) {
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == '0' {
		return
	}
	visited[i][j] = true
	*found = true
	numIslandsDFS(grid, visited, m, n, i-1, j, found)
	numIslandsDFS(grid, visited, m, n, i+1, j, found)
	numIslandsDFS(grid, visited, m, n, i, j-1, found)
	numIslandsDFS(grid, visited, m, n, i, j+1, found)
}
