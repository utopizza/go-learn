package leetcode

func trap(height []int) int {
	sum := 0
	leftMax := leftHighest(height)
	rightMax := rightHighest(height)
	for i := 0; i < len(height); i++ {
		water := min(leftMax[i], rightMax[i]) - height[i]
		if water > 0 {
			sum += water
		}
	}
	return sum
}

func leftHighest(height []int) []int {
	m := 0
	res := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if height[i] >= m {
			m = height[i]
		}
		res[i] = m
	}
	return res
}

func rightHighest(height []int) []int {
	m := 0
	res := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] >= m {
			m = height[i]
		}
		res[i] = m
	}
	return res
}
