package tree

import "github.com/utopizza/go-learn/leetcode/utils"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftMin := minDepth(root.Left)
	rightMin := minDepth(root.Right)

	if leftMin > 0 && rightMin > 0 {
		return utils.Min(leftMin, rightMin) + 1
	} else if leftMin == 0 {
		return rightMin + 1
	} else {
		return leftMin + 1
	}
}
