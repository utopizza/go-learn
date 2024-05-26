package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	combinationSum2Recursive(candidates, target, 0, 0, []int{}, &result)
	return result
}

func combinationSum2Recursive(candidates []int, target int, index int, sum int, arr []int, result *[][]int) {
	if sum > target {
		return
	} else if sum == target {
		*result = append(*result, append([]int(nil), arr...))
		return
	} else {
		for i := index + 1; i < len(candidates); i++ {
			combinationSum2Recursive(candidates, target, i, sum+candidates[i], append(arr, candidates[i]), result)
		}
	}
}
