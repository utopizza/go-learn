package array

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for k := 0; k < len(nums)-1; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		target := 0 - nums[k]
		for lo, hi := k+1, len(nums)-1; lo < hi; {
			if lo > k+1 && nums[lo] == nums[lo-1] {
				lo++
				continue
			}
			if lo >= hi {
				break
			}
			sum := nums[lo] + nums[hi]
			if sum == target {
				result = append(result, []int{nums[k], nums[lo], nums[hi]})
				lo++
			} else if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return result
}
