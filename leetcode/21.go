package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{0, nil}
	tail := temp
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			tail.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val <= list2.Val {
				tail.Next = list1
				list1 = list1.Next
			} else {
				tail.Next = list2
				list2 = list2.Next
			}
		}
		tail = tail.Next
	}
	return temp.Next
}
