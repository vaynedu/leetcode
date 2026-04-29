package interview

import "testing"

func TestNumIslandsSingleProblemCases(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "单个岛屿",
			grid: byteGrid([]string{
				"11110",
				"11010",
				"11000",
				"00000",
			}),
			want: 1,
		},
		{
			name: "多个岛屿",
			grid: byteGrid([]string{
				"11000",
				"11000",
				"00100",
				"00011",
			}),
			want: 3,
		},
		{
			name: "对角不连通",
			grid: byteGrid([]string{
				"10",
				"01",
			}),
			want: 2,
		},
		{
			name: "全水",
			grid: byteGrid([]string{
				"000",
				"000",
			}),
			want: 0,
		},
		{
			name: "空网格",
			grid: [][]byte{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(cloneByteGrid(tt.grid))
			if got != tt.want {
				t.Fatalf("numIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNumIslandsByUFSingleProblemCases(t *testing.T) {
	grid := byteGrid([]string{
		"11000",
		"11000",
		"00100",
		"00011",
	})
	if got := numIslandsByUF(cloneByteGrid(grid)); got != 3 {
		t.Fatalf("numIslandsByUF() = %d, want 3", got)
	}
}

func byteGrid(rows []string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, row := range rows {
		grid[i] = []byte(row)
	}
	return grid
}

func cloneByteGrid(grid [][]byte) [][]byte {
	if len(grid) == 0 {
		return [][]byte{}
	}
	cp := make([][]byte, len(grid))
	for i := range grid {
		cp[i] = append([]byte(nil), grid[i]...)
	}
	return cp
}
