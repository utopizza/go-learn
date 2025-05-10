package dp

func partition(s string) [][]string {
	// 动态规划预处理s[i:j]是否回文串
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	// 回溯法切割s搜索解空间
	var solutions [][]string
	var solution []string
	partitionBackTrack(s, 0, solution, &solutions, dp)
	return solutions
}

func partitionBackTrack(s string, i int, solution []string, solutions *[][]string, dp [][]bool) {
	if i == len(s) {
		*solutions = append(*solutions, append([]string{}, solution...))
		return
	}
	for j := i; j < len(s); j++ {
		if dp[i][j] {
			partitionBackTrack(s, j+1, append(solution, s[i:j+1]), solutions, dp)
		}
	}
}
