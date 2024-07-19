package linklist

func partition(head *ListNode, x int) *ListNode {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	p, q := l1, l2
	for r := head; r != nil; r = r.Next {
		if r.Val < x {
			p.Next = r
			p = p.Next
		} else {
			q.Next = r
			q = q.Next
		}
	}
	q.Next = nil
	p.Next = l2.Next
	return l1.Next
}
