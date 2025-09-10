package leetcode

// 994.腐烂的橘子

import (
	"testing"
)

/**
在给定的 m x n 网格
 grid 中，每个单元格可以有以下三个值之一：


 值 0 代表空单元格；
 值 1 代表新鲜橘子；
 值 2 代表腐烂的橘子。


 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。



 示例 1：




输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4


 示例 2：


输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。


 示例 3：


输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。




 提示：


 m == grid.length
 n == grid[i].length
 1 <= m, n <= 10
 grid[i][j] 仅为 0、1 或 2


 👍 1060 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	// 定义四个方向：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 队列存储腐烂橘子的坐标
	queue := [][]int{}

	// 统计新鲜橘子数量
	freshCount := 0

	// 初始化：找到所有腐烂橘子和新鲜橘子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if freshCount == 0 {
		return 0
	}

	// BFS过程
	minutes := 0
	for len(queue) > 0 && freshCount > 0 {
		// 当前这一分钟要处理的腐烂橘子数量
		currentRotten := len(queue)

		// 处理当前所有腐烂橘子同时扩散
		for i := 0; i < currentRotten; i++ {
			// 取出一个腐烂橘子
			rotten := queue[0]
			queue = queue[1:]

			// 向四个方向扩散
			for _, dir := range directions {
				newRow, newCol := rotten[0]+dir[0], rotten[1]+dir[1]

				// 检查边界和是否为新鲜橘子
				if newRow >= 0 && newRow < rows &&
					newCol >= 0 && newCol < cols &&
					grid[newRow][newCol] == 1 {

					// 腐烂该橘子
					grid[newRow][newCol] = 2
					freshCount--

					// 加入队列，下一分钟会继续扩散
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}

		// 过完一分钟
		minutes++
	}

	// 如果还有新鲜橘子，说明无法全部腐烂
	if freshCount > 0 {
		return -1
	}

	return minutes
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region end(Prohibit modification and deletion)

func TestRottingOranges(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "示例2",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name:     "示例3",
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
		{
			name:     "全部新鲜",
			grid:     [][]int{{1, 1, 1}},
			expected: -1,
		},
		{
			name:     "全部腐烂",
			grid:     [][]int{{2, 2, 2}},
			expected: 0,
		},
		{
			name:     "全部空格",
			grid:     [][]int{{0, 0, 0}},
			expected: 0,
		},
		{
			name: "复杂情况",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{0, 1, 2},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建grid的副本，避免修改原始数据
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := orangesRotting(gridCopy)
			if result != tt.expected {
				t.Errorf("orangesRotting(%v) = %d, expected %d", tt.grid, result, tt.expected)
			}
		})
	}
}
