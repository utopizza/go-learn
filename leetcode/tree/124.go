package tree

import (
	"github.com/utopizza/go-learn/leetcode/utils"
	"math"
)

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	maxGain(root, &maxSum)
	return maxSum
}

func maxGain(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftGain := utils.Max(maxGain(node.Left, maxSum), 0)
	rightGain := utils.Max(maxGain(node.Right, maxSum), 0)

	priceNewPath := node.Val + leftGain + rightGain
	*maxSum = utils.Max(*maxSum, priceNewPath)

	return node.Val + utils.Max(leftGain, rightGain)
}
