package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursive(root.Left, root.Right)
}

func isSymmetricRecursive(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricRecursive(p.Left, q.Right) && isSymmetricRecursive(p.Right, q.Left)
}
