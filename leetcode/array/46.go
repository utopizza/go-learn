package array

func permute(nums []int) [][]int {
	var solution []int
	var solutions [][]int
	visited := make([]bool, len(nums))
	permuteBackTrack(nums, visited, solution, &solutions)
	return solutions
}

func permuteBackTrack(nums []int, visited []bool, solution []int, solutions *[][]int) {
	if len(solution) == len(nums) {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		solution = append(solution, nums[i])
		permuteBackTrack(nums, visited, solution, solutions)
		solution = solution[:len(solution)-1]
		visited[i] = false
	}
}
