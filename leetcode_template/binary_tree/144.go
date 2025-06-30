package binary_tree

func preorderTraversal(root *TreeNode) []int {
	var nodes []int
	preorderTraversalRecursive(root, &nodes)
	return nodes
}

func preorderTraversalRecursive(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root.Val)
	preorderTraversalRecursive(root.Left, nodes)
	preorderTraversalRecursive(root.Right, nodes)
}
