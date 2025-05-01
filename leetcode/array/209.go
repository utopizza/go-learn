package array

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minLen := math.MaxInt
	for lo, hi := 0, 0; hi < len(nums); hi++ {
		sum += nums[hi]
		if sum < target {
			continue
		}
		for sum >= target {
			minLen = min(minLen, hi-lo+1)
			sum -= nums[lo]
			lo++
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
