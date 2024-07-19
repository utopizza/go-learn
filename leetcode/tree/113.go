package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var currPath []int
	pathSumRecursive(root, targetSum, currPath, &result)
	return result
}

func pathSumRecursive(node *TreeNode, targetSum int, currentPath []int, solutions *[][]int) {
	if node == nil {
		return
	}

	currentPath = append(currentPath, node.Val)
	defer func() {
		currentPath = currentPath[:len(currentPath)-1]
	}()

	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		*solutions = append(*solutions, append([]int{}, currentPath...))
		return
	}

	pathSumRecursive(node.Left, targetSum-node.Val, currentPath, solutions)
	pathSumRecursive(node.Right, targetSum-node.Val, currentPath, solutions)
}
