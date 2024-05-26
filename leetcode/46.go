package leetcode

func permute(nums []int) [][]int {
	var solutions [][]int
	used := make(map[int]bool)
	permuteRecursive(nums, []int{}, &solutions, used)
	return solutions
}

func permuteRecursive(nums []int, arr []int, solutions *[][]int, used map[int]bool) {
	if len(arr) == len(nums) {
		*solutions = append(*solutions, append([]int{}, arr...))
	} else {
		for _, n := range nums {
			if !used[n] {
				used[n] = true
				permuteRecursive(nums, append(arr, n), solutions, used)
				delete(used, n)
			}
		}
	}
}
