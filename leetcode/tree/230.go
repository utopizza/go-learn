package tree

func kthSmallest(root *TreeNode, k int) int {
	arr := inorderTraversal(root)
	if k-1 >= len(arr) {
		return -1
	}
	return arr[k-1]
}
