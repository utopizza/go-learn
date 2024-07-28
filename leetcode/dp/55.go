package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func canJumpGreedy(nums []int) bool {
	far := nums[0]
	for i := range nums {
		if far >= i {
			far = utils.Max(far, i+nums[i])
		} else {
			break
		}
	}
	return far >= len(nums)-1
}
