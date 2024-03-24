package leetcode

func generateParenthesis(n int) []string {
	var results []string
	backtrack(n, n, "", &results)
	return results
}

func backtrack(left, right int, curStr string, results *[]string) {
	if left == 0 && right == 0 {
		*results = append(*results, curStr)
		return
	}
	if left > 0 {
		backtrack(left-1, right, curStr+"(", results)
	}
	if right > left {
		backtrack(left, right-1, curStr+")", results)
	}
}
