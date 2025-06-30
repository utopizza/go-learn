package backtracking

func permute(nums []int) [][]int {
	var solution []int
	var solutions [][]int
	visited := make(map[int]bool)
	permuteBacktrack(nums, visited, solution, &solutions)
	return solutions
}

func permuteBacktrack(nums []int, visited map[int]bool, solution []int, solutions *[][]int) {
	if len(solution) == len(nums) {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		solution = append(solution, nums[i])
		visited[nums[i]] = true
		permuteBacktrack(nums, visited, solution, solutions)
		visited[nums[i]] = false
		solution = solution[:len(solution)-1]
	}
}
