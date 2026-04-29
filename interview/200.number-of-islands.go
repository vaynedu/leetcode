package interview

// ================================================================
// 200. 岛屿数量
// ================================================================

// numIslands DFS 解法
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	count := 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, d := range dirs {
			dfs(i+d[0], j+d[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// ================================================================
// 并查集版本（备选）
// ================================================================

// UnionFind 并查集
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			}
		}
	}
	return &UnionFind{parent: parent, rank: rank, count: count}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		if uf.rank[rootX] == uf.rank[rootY] {
			uf.rank[rootX]++
		}
		uf.count--
	}
}

func numIslandsByUF(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	uf := newUnionFind(grid)
	dirs := [][]int{{1, 0}, {0, 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
						uf.union(i*n+j, ni*n+nj)
					}
				}
			}
		}
	}
	return uf.count
}
