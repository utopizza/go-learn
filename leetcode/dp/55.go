package dp

func canJump(nums []int) bool {
	far := nums[0]
	for i := 1; i < len(nums); i++ {
		if i <= far {
			far = max(far, i+nums[i])
		}
	}
	return far >= len(nums)-1
}
