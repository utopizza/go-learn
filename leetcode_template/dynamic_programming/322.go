package dynamic_programming

import "math"

func coinChange(coins []int, amount int) int {
	// dp[i]=min{dp[i-coin[j]]}+1
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ { // 尝试用第j枚作为凑i的最后一枚硬币
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
