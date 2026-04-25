package interview

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 3},
		{"4", 4, 5},
		{"5", 5, 8},
		{"10", 10, 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
