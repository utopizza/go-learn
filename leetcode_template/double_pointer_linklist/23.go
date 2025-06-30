package double_pointer_linklist

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l <= 1 {
		return lists[0]
	}
	left := mergeKLists(lists[0 : l/2])
	right := mergeKLists(lists[l/2:])
	return mergeTwoLists(left, right)
}
