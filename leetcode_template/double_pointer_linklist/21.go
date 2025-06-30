package double_pointer_linklist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{-1, nil}

	tail := preHead
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			tail.Next = list1
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
		tail.Next = nil
	}

	return preHead.Next
}
