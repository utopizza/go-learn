package tree

func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	l := len(result)
	for i := 0; i < l/2; i++ {
		result[i], result[l-1-i] = result[l-1-i], result[i]
	}
	return result
}
