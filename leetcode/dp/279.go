package dp

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minN := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minN = min(minN, dp[i-j*j])
		}
		dp[i] = minN + 1
	}
	return dp[n]
}
