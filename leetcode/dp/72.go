package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}

	// dp矩阵
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// dp矩阵边界初始化
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	// dp矩阵填充
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			add := dp[i-1][j] + 1 // word1[:i-1]末尾插入1字符变成word2[:j]
			del := dp[i][j-1] + 1 // word2[:j-1]末尾插入1字符变成word1[:i]
			var replace int       // word1[i]与word[j]相等，不用操作；否则替换其中一个
			if word1[i-1] != word2[j-1] {
				replace = dp[i-1][j-1] + 1
			} else {
				replace = dp[i-1][j-1]
			}
			dp[i][j] = utils.Min(add, utils.Min(del, replace))
		}
	}

	return dp[m][n]
}
