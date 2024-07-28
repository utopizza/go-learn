package dp

import (
	"math"
)

func minCut(s string) int {
	n := len(s)

	// dp[i][j]表示s[i..j]是否回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
		}
	}

	// dp2[i]表示s[0..i]分割回文串的最小次数
	dp2 := make([]int, n)
	for i := 0; i < n; i++ {
		if dp[0][i] {
			continue
		}
		dp2[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if dp[j+1][i] && dp2[j]+1 < dp2[i] {
				dp2[i] = dp2[j] + 1
			}
		}
	}

	return dp2[n-1]
}
