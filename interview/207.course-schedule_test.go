package interview

import "testing"

func TestCanFinish(t *testing.T) {
    tests := []struct {
        name string
        num int
        pre [][]int
        want bool
    }{
        {"可完成", 4, [][]int{{1,0},{2,0},{3,1},{3,2}}, true},
        {"有环", 2, [][]int{{1,0},{0,1}}, false},
        {"单课程", 1, [][]int{}, true},
        {"线性依赖", 3, [][]int{{1,0},{2,1}}, true},
        {"多入度", 4, [][]int{{1,0},{2,0},{3,0}}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := canFinish(tt.num, tt.pre)
            if got != tt.want {
                t.Errorf("canFinish(%d, %v) = %v, want %v", tt.num, tt.pre, got, tt.want)
            }
        })
    }
}
