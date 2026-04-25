package interview

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0", 0, 1},
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 5},
		{"4", 4, 14},
		{"5", 5, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numTrees(tt.n)
			if got != tt.want {
				t.Errorf("numTrees(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
