package array

import "fmt"

func summaryRanges(nums []int) []string {
	var res []string
	n := len(nums)
	for lo, hi := 0, 0; hi < n; hi++ {
		if hi < n-1 && nums[hi]+1 == nums[hi+1] {
			continue
		}
		if lo == hi {
			res = append(res, fmt.Sprintf("%d", nums[lo]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[lo], nums[hi]))
		}
		lo = hi + 1
	}
	return res
}
