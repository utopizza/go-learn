package leetcode

import "math"

func jump(nums []int) int {
	// 动态规划
	dp := make([]int, len(nums))
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(dp)-1]
}

func jumpGreedy(nums []int) int {
	// 用贪心
	end := 0
	step := 0
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			end = farthest
			step++
		}
	}
	return step
}
