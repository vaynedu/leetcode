package interview

import "testing"

func TestLadderLength(t *testing.T) {
    tests := []struct {
        name string
        begin string
        end string
        list []string
        want int
    }{
        {"基本", "hit", "cog", []string{"hot","dot","dog","lot","log","cog"}, 5},
        {"无路径", "hit", "cog", []string{"hot","dot","dog","lot","log"}, 0},
        {"直接", "a", "c", []string{"a","b","c"}, 2},
        {"单步", "a", "b", []string{"a","b"}, 2},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ladderLength(tt.begin, tt.end, tt.list)
            if got != tt.want {
                t.Errorf("ladderLength(%s, %s, %v) = %d, want %d", tt.begin, tt.end, tt.list, got, tt.want)
            }
        })
    }
}
