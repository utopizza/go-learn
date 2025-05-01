package array

import "math"

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(dp[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	minSum := math.MaxInt
	for _, v := range dp[len(dp)-1] {
		minSum = min(minSum, v)
	}
	return minSum
}
