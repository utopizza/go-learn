package linklist

func deleteNode(node *ListNode) {
	var pre *ListNode
	for p := node; p.Next != nil; p = p.Next {
		p.Val = p.Next.Val
		if p.Next.Next == nil {
			pre = p
		}
	}
	pre.Next = nil
}
