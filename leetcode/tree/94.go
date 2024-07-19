package tree

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorderRecursive(root, &result)
	return result
}

func inorderRecursive(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderRecursive(node.Left, result)
	*result = append(*result, node.Val)
	inorderRecursive(node.Right, result)
}
