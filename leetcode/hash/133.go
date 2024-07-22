package hash

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// copy node
	nodeMap := make(map[*Node]*Node)
	visited := make(map[*Node]bool)
	cloneGraphCopyNode(node, visited, nodeMap)

	// copy edges
	for ori, newNode := range nodeMap {
		oriNeighbors := ori.Neighbors
		var newNeighbors []*Node
		for _, o := range oriNeighbors {
			c := nodeMap[o]
			newNeighbors = append(newNeighbors, c)
		}
		newNode.Neighbors = newNeighbors
	}

	return nodeMap[node]
}

func cloneGraphCopyNode(node *Node, visited map[*Node]bool, nodeMap map[*Node]*Node) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true
	nodeMap[node] = &Node{Val: node.Val, Neighbors: nil}
	for _, n := range node.Neighbors {
		cloneGraphCopyNode(n, visited, nodeMap)
	}
}
