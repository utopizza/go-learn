package array

func maxArea(height []int) int {
	area := 0
	lo, hi := 0, len(height)-1
	for lo < hi {
		if height[lo] < height[hi] {
			area = max(area, height[lo]*(hi-lo))
			lo++
		} else {
			area = max(area, height[hi]*(hi-lo))
			hi--
		}
	}
	return area
}
