package double_pointer_linklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{-1, head}

	// fast先走n步
	slow, fast := preHead, preHead
	for i := n; i >= 0; i-- {
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	// slow和fast一起走，直至fast到队尾
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除slow的下一个节点
	if slow != nil {
		slow.Next = slow.Next.Next
	}

	return preHead.Next
}
