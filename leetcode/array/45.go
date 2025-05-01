package array

import "math"

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		minJump := math.MaxInt
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				minJump = min(minJump, dp[j]+1)
			}
		}
		dp[i] = minJump
	}
	return dp[len(nums)-1]
}
