package linklist

func insertionSortList(head *ListNode) *ListNode {
	tempHead := &ListNode{0, head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = cur
		} else {
			prev := tempHead
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return tempHead.Next
}
