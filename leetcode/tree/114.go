package tree

func flatten(root *TreeNode) {
	var arr []*TreeNode
	flattenPreOrderTraversal(root, &arr)
	for i := 0; i < len(arr)-1; i++ {
		arr[i].Left, arr[i].Right = nil, arr[i+1]
	}
}

func flattenPreOrderTraversal(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	*arr = append(*arr, root)
	flattenPreOrderTraversal(root.Left, arr)
	flattenPreOrderTraversal(root.Right, arr)
}
