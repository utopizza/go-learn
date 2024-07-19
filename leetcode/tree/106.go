package tree

func buildTreeFromInorderAndPostorder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	leftCnt := len(inorder[:i])
	root.Left = buildTreeFromInorderAndPostorder(inorder[:i], postorder[:leftCnt])
	root.Right = buildTreeFromInorderAndPostorder(inorder[i+1:], postorder[leftCnt:len(postorder)-1])
	return root
}
