package array

func trap(height []int) int {
	var sum int

	// 每个柱子左边的最高高度
	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// 每个柱子右边的最高高度
	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// 每个柱子上面的最大积水 = min(最左最大高度，最右最大高度) - 当前柱子本身高度
	for i := range height {
		cur := min(leftMax[i], rightMax[i]) - height[i]
		if cur >= 0 {
			sum += cur
		}
	}
	return sum
}
