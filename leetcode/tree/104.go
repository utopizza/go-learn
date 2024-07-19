package tree

import "github.com/utopizza/go-learn/leetcode/utils"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return utils.Max(leftMax, rightMax) + 1
}
