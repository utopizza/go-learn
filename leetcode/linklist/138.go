package linklist

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	for p := head; p != nil; p = p.Next {
		m[p] = &Node{Val: p.Val}
	}

	for p := head; p != nil; p = p.Next {
		if p.Next != nil {
			m[p].Next = m[p.Next]
		}
		if p.Random != nil {
			m[p].Random = m[p.Random]
		}
	}

	return m[head]
}
