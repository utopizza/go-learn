package dp

func generateParenthesis(n int) []string {
	var solutions []string
	generateParenthesisBacktrack(n, 0, 0, "", &solutions)
	return solutions
}

func generateParenthesisBacktrack(n int, l int, r int, solution string, solutions *[]string) {
	if l+r == n {
		*solutions = append(*solutions, solution)
		return
	}
	if l < n {
		generateParenthesisBacktrack(n, l+1, r, solution+"(", solutions)
	}
	if r < l {
		generateParenthesisBacktrack(n, l, r+1, solution+")", solutions)
	}
}
