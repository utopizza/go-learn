package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 神经病的lt用例
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// 初始化
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	// 推算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			top := dp[i-1][j]
			if obstacleGrid[i-1][j] == 1 {
				top = 0
			}
			left := dp[i][j-1]
			if obstacleGrid[i][j-1] == 1 {
				left = 0
			}
			dp[i][j] = top + left
		}
	}

	return dp[m-1][n-1]
}
