package interview

import "testing"

func TestIsBipartite(t *testing.T) {
    tests := []struct {
        name string
        graph [][]int
        want bool
    }{
        {"环4", [][]int{{1,2,3},{0,2},{0,1},{0}}, false}, // 3节点环
        {"偶数环", [][]int{{1,3},{0,2},{1,3},{0,2}}, true},
        {"单节点", [][]int{{}}, true},
        {"两个节点", [][]int{{1},{0}}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := isBipartite(tt.graph)
            if got != tt.want {
                t.Errorf("isBipartite() = %v, want %v", got, tt.want)
            }
        })
    }
}
