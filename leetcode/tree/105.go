package tree

func buildTreeFromPreorderAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	leftCnt := len(inorder[:i])
	root.Left = buildTreeFromPreorderAndInorder(preorder[1:leftCnt+1], inorder[:i])
	root.Right = buildTreeFromPreorderAndInorder(preorder[leftCnt+1:], inorder[i+1:])
	return root
}
