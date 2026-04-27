package interview

import "testing"

func TestExist(t *testing.T) {
    board := [][]byte{
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'},
    }
    tests := []struct {
        name string
        board [][]byte
        word string
        want bool
    }{
        {"存在", board, "ABCCED", true},
        {"存在2", board, "SEE", true},
        {"不存在", board, "ABCB", false},
        {"单字符", [][]byte{{'a'}}, "a", true},
        {"不存在2", board, "ABFB", false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := exist(tt.board, tt.word)
            if got != tt.want {
                t.Errorf("exist() = %v, want %v", got, tt.want)
            }
        })
    }
}
