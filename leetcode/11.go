package leetcode

func maxArea(height []int) int {
	var area int
	for i, j := 0, len(height)-1; i < j; {
		area = max(area, (j-i)*min(height[i], height[j]))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}
