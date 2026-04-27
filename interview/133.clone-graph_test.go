package interview

import "testing"

func TestCloneGraph(t *testing.T) {
	// 构造图: 1--2, 1--3, 2--4, 3--4
	n1 := &GraphNode{Val: 1}
	n2 := &GraphNode{Val: 2}
	n3 := &GraphNode{Val: 3}
	n4 := &GraphNode{Val: 4}
	n1.Neighbors = []*GraphNode{n2, n3}
	n2.Neighbors = []*GraphNode{n1, n4}
	n3.Neighbors = []*GraphNode{n1, n4}
	n4.Neighbors = []*GraphNode{n2, n3}

	cloned := cloneGraph(n1)
	if cloned == nil || cloned.Val != 1 {
		t.Errorf("cloneGraph() val = %d, want 1", cloned.Val)
	}
	if len(cloned.Neighbors) != 2 {
		t.Errorf("cloned.Neighbors len = %d, want 2", len(cloned.Neighbors))
	}
	if cloned == n1 {
		t.Errorf("cloneGraph should return a new object")
	}
}
