package array

func subsets(nums []int) [][]int {
	var solutions [][]int
	var solution []int
	subsetsBackTrack(nums, 0, solution, &solutions)
	return solutions
}

func subsetsBackTrack(nums []int, index int, solution []int, solutions *[][]int) {
	*solutions = append(*solutions, append([]int{}, solution...))
	for i := index; i < len(nums); i++ {
		solution = append(solution, nums[i])
		subsetsBackTrack(nums, i+1, solution, solutions)
		solution = solution[:len(solution)-1]
	}
}
