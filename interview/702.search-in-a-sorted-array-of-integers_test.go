package interview

import "testing"

type MockReader struct {
    arr []int
}

func (m MockReader) Get(i int) int {
    if i < 0 || i >= len(m.arr) {
        return 1<<31 - 1
    }
    return m.arr[i]
}

func TestSearch(t *testing.T) {
    reader := MockReader{arr: []int{-1, 0, 3, 5, 9, 12}}
    tests := []struct {
        name string
        target int
        want int
    }{
        {"存在", 9, 4},
        {"不存在", 2, -1},
        {"第一个", -1, 0},
        {"边界", 12, 5},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := search(reader, tt.target)
            if got != tt.want {
                t.Errorf("search() = %d, want %d", got, tt.want)
            }
        })
    }
}
