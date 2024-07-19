package linklist

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddTempHead := &ListNode{}
	evenTempHead := &ListNode{}
	oddCur, evenCur, cur := oddTempHead, evenTempHead, head

	i := 0
	for cur != nil {
		i++
		if i%2 == 1 {
			oddCur.Next = cur
			oddCur = oddCur.Next
			cur = cur.Next
			oddCur.Next = nil
		} else {
			evenCur.Next = cur
			evenCur = evenCur.Next
			cur = cur.Next
			evenCur.Next = nil
		}
	}

	oddCur.Next = evenTempHead.Next
	return oddTempHead.Next
}
