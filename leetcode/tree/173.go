package tree

type BSTIterator struct {
	index    int
	innerArr []int
}

func Constructor(root *TreeNode) BSTIterator {
	arr := inorderTraversal(root)
	return BSTIterator{
		index:    0,
		innerArr: arr,
	}
}

func (i *BSTIterator) Next() int {
	ret := i.innerArr[i.index]
	i.index++
	return ret
}

func (i *BSTIterator) HasNext() bool {
	return i.index < len(i.innerArr)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
