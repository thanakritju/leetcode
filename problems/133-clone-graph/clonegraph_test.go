package clonegraph

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		{
			name: "Empty graph",
			node: nil,
		},
		{
			name: "Graph with two connected nodes",
			node: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node1.Neighbors = []*Node{node2}
				node2.Neighbors = []*Node{node1}
				return node1
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := cloneGraph(tt.node)

			// Check if the clone matches the original structure
			if !reflect.DeepEqual(cloned, tt.node) {
				t.Errorf("cloneGraph() = %v, want %v", cloned, tt.node)
			}

			// Modify the original graph to ensure deep copy
			if tt.node != nil && len(tt.node.Neighbors) > 0 {
				tt.node.Neighbors[0].Val = 999
				if reflect.DeepEqual(cloned, tt.node) {
					t.Errorf("cloneGraph() did not produce a deep copy")
				}
			}
		})
	}
}
