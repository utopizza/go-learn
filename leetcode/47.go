package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var solutions [][]int
	used := make(map[int]bool)
	sort.Ints(nums)
	permuteUniqueRecursive(nums, []int{}, &solutions, used)
	return solutions
}

func permuteUniqueRecursive(nums []int, arr []int, solutions *[][]int, used map[int]bool) {
	if len(arr) == len(nums) {
		*solutions = append(*solutions, append([]int{}, arr...))
	} else {
		for i := 0; i < len(nums); i++ {
			if i > 0 && used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				permuteUniqueRecursive(nums, append(arr, nums[i]), solutions, used)
				delete(used, i)
			}
		}
	}
}
