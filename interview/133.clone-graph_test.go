package interview

import "testing"

func TestCloneGraph(t *testing.T) {
    // 构造图: 1--2, 1--3, 2--4, 3--4
    n1 := &Node{Val: 1}
    n2 := &Node{Val: 2}
    n3 := &Node{Val: 3}
    n4 := &Node{Val: 4}
    n1.Neighbors = []*Node{n2, n3}
    n2.Neighbors = []*Node{n1, n4}
    n3.Neighbors = []*Node{n1, n4}
    n4.Neighbors = []*Node{n2, n3}

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
