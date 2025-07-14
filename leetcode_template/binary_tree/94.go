package binary_tree

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 有左孩子，一直压入栈，模拟程序栈dfs
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 左子树访问完，开始退「栈」，弹出栈顶
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 访问根，实现中序访问
		res = append(res, root.Val)
		// 访问右子树
		root = root.Right
	}
	return res
}
