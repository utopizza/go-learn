package tree

func sumNumbers(root *TreeNode) int {
	var cur, sum int
	sumNumbersRecur(root, &cur, &sum)
	return sum
}

func sumNumbersRecur(node *TreeNode, cur *int, sum *int) {
	if node == nil {
		return
	}

	*cur = *cur*10 + node.Val
	defer func() {
		*cur = *cur / 10
	}()

	if node.Left == nil && node.Right == nil {
		*sum = *sum + *cur
		return
	}

	sumNumbersRecur(node.Left, cur, sum)
	sumNumbersRecur(node.Right, cur, sum)
}
