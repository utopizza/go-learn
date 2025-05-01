package array

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var solution []int
	var solutions [][]int
	combinationSumBackTrack(candidates, target, 0, 0, solution, &solutions)
	return solutions
}

func combinationSumBackTrack(candidates []int, target int, index int, sum int, solution []int, solutions *[][]int) {
	if sum == target {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		solution = append(solution, candidates[i])
		combinationSumBackTrack(candidates, target, i, sum+candidates[i], solution, solutions)
		solution = solution[:len(solution)-1]
	}
}
