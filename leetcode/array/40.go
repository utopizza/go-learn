package array

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var solution []int
	var solutions [][]int
	combinationSum2BackTrack(candidates, target, 0, 0, solution, &solutions)
	return solutions
}

func combinationSum2BackTrack(candidates []int, target int, index int, sum int, solution []int, solutions *[][]int) {
	if sum == target {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		solution = append(solution, candidates[i])
		combinationSum2BackTrack(candidates, target, i+1, sum+candidates[i], solution, solutions)
		solution = solution[:len(solution)-1]
	}
}
