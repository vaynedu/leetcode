package interview

import "testing"

func TestFindOrder(t *testing.T) {
    tests := []struct {
        name string
        num int
        pre [][]int
        wantLen int
    }{
        {"可完成", 4, [][]int{{1,0},{2,0},{3,1},{3,2}}, 4},
        {"有环", 2, [][]int{{1,0},{0,1}}, 0},
        {"线性", 3, [][]int{{1,0},{2,1}}, 3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findOrder(tt.num, tt.pre)
            if len(got) != tt.wantLen {
                t.Errorf("findOrder(%d, %v) len=%d, want %d", tt.num, tt.pre, len(got), tt.wantLen)
            }
        })
    }
}
