package tree

import "github.com/utopizza/go-learn/leetcode/utils"

func isBalanced(root *TreeNode) bool {
	result, _ := isBalancedRecursive(root)
	return result
}

func isBalancedRecursive(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	isLeftBalanced, leftMaxDepth := isBalancedRecursive(node.Left)
	isRightBalanced, rightMaxDepth := isBalancedRecursive(node.Right)
	if !isLeftBalanced || !isRightBalanced {
		return false, -1
	}
	if utils.Abs(leftMaxDepth, rightMaxDepth) > 1 {
		return false, -1
	}
	return true, utils.Max(leftMaxDepth, rightMaxDepth) + 1
}
