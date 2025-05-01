package array

func combinationSumV3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var solution []int
	var solutions [][]int
	combinationSumV3BackTrack(nums, k, n, 0, solution, &solutions)
	return solutions
}

func combinationSumV3BackTrack(nums []int, k, n, i int, solution []int, solutions *[][]int) {
	if n == 0 && k == 0 {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}
	if n < 0 || k < 0 {
		return
	}
	for j := i; j < len(nums); j++ {
		solution = append(solution, nums[j])
		combinationSumV3BackTrack(nums, k-1, n-nums[j], j+1, solution, solutions)
		solution = solution[:len(solution)-1]
	}
}
