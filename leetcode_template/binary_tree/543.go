package binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	var maxD int
	diameterOfBinaryTreeRecursive(root, &maxD)
	return maxD - 1
}

func diameterOfBinaryTreeRecursive(root *TreeNode, maxD *int) int {
	if root == nil {
		return 0
	}
	leftD := diameterOfBinaryTreeRecursive(root.Left, maxD)   // 左子树深度
	rightD := diameterOfBinaryTreeRecursive(root.Right, maxD) // 右子树深度
	*maxD = max(*maxD, leftD+rightD+1)                        // 最大直径
	return max(leftD, rightD) + 1                             // 返回最大深度
}
