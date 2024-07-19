package tree

func isValidBST(root *TreeNode) bool {
	arr := inorderTraversal(root)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
