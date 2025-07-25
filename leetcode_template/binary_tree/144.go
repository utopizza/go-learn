package binary_tree

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 先根访问，访问完入栈
			res = append(res, root.Val)
			stack = append(stack, root)
			// 访问左孩子
			root = root.Left
		}
		// 左孩子访问完，开始退栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 访问右孩子
		root = root.Right
	}
	return res
}
