package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)

	var cloneNode func(*Node) *Node
	cloneNode = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if clonedNode, found := visited[n]; found {
			return clonedNode
		}

		clonedNode := &Node{Val: n.Val}
		visited[n] = clonedNode

		for _, neighbor := range n.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, cloneNode(neighbor))
		}

		return clonedNode
	}

	return cloneNode(node)
}
