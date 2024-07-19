package linklist

func removeElements(head *ListNode, val int) *ListNode {
	tempHead := &ListNode{}
	cur := tempHead
	for head != nil {
		if head.Val != val {
			cur.Next = head
			head = head.Next
			cur = cur.Next
			cur.Next = nil
		} else {
			head = head.Next
		}
	}
	return tempHead.Next
}
