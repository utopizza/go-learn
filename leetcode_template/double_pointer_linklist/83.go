package double_pointer_linklist

func deleteDuplicates(head *ListNode) *ListNode {
	for p := head; p != nil; {
		if p.Next != nil && p.Next.Val == p.Val {
			node := p.Next
			p.Next = p.Next.Next
			node.Next = nil
		} else {
			p = p.Next
		}
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
