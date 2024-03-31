package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var arr []int
	var result [][]int
	combinationSumRecursive(candidates, target, 0, 0, arr, &result)
	return result
}

func combinationSumRecursive(candidates []int, target int, index int, sum int, arr []int, result *[][]int) {
	if sum > target {
		return
	} else if sum == target {
		*result = append(*result, append([]int(nil), arr...)) // 切记这里要新建一个数组
		return
	} else {
		for i := index; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			combinationSumRecursive(candidates, target, i, sum+candidates[i], arr, result)
			arr = arr[:len(arr)-1]
		}
	}
}
