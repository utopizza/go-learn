package tree

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	var lastVisit *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == lastVisit { // top无右子树了，或者top的右子树已经访问完了
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			lastVisit = top
		} else { // top的右子树压入
			root = top.Right
		}
	}
	return result
}
