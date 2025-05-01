package array

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var solution []int
	var solutions [][]int
	visited := make([]bool, len(nums))
	permuteUniqueBackTrack(nums, visited, solution, &solutions)
	return solutions
}

func permuteUniqueBackTrack(nums []int, visited []bool, solution []int, solutions *[][]int) {
	if len(solution) == len(nums) {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] || visited[i] {
			continue
		}
		visited[i] = true
		solution = append(solution, nums[i])
		permuteUniqueBackTrack(nums, visited, solution, solutions)
		solution = solution[:len(solution)-1]
		visited[i] = false
	}
}
