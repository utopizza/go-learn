package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	for _, list := range lists {
		head = mergeTwoLists(head, list)
	}
	return head
}
