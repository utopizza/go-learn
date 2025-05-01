package array

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var solutions [][]int
	var solution []int
	visited := make([]bool, len(nums))
	subsetsWithDupBackTrack(nums, 0, solution, &solutions, visited)
	return solutions
}

func subsetsWithDupBackTrack(nums []int, index int, solution []int, solutions *[][]int, visited []bool) {
	*solutions = append(*solutions, append([]int{}, solution...))
	for i := index + 1; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] { // 若发现没有选择上一个数，且当前数字与上一个数相同，则可以跳过
			continue
		}
		solution = append(solution, nums[i])
		visited[i] = true
		subsetsWithDupBackTrack(nums, i, solution, solutions, visited)
		visited[i] = false
		solution = solution[:len(solution)-1]
	}
}
