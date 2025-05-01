package array

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	lo, hi := 0, n-1
	for {
		mid := lo + (hi-lo)/2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			lo = mid + 1 // 往右走，丢弃左边
		} else {
			hi = mid - 1 // 往左走，丢弃右边
		}
	}
}
