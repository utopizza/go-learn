package dp

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
