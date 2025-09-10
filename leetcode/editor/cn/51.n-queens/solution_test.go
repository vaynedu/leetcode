package leetcode

// 51.N 皇后

import (
	"testing"
)

/**
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

 n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。



 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。





 示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。


 示例 2：


输入：n = 1
输出：[["Q"]]




 提示：


 1 <= n <= 9


 👍 2338 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var result [][]string
	// 用于记录皇后的位置，queens[i]表示第i行皇后所在的列
	queens := make([]int, n)

	// 用于记录已被占用的列
	cols := make([]bool, n)
	// 用于记录主对角线（行-列）
	diag1 := make([]bool, 2*n-1)
	// 用于记录副对角线（行+列）
	diag2 := make([]bool, 2*n-1)

	// 回溯函数
	var backtrack func(row int)
	backtrack = func(row int) {
		// 如果所有行都已放置皇后，说明找到一个解
		if row == n {
			// 将当前解转换为题目要求的格式
			result = append(result, generateBoard(queens, n))
			return
		}

		// 在当前行尝试每一列
		for col := 0; col < n; col++ {
			// 检查当前位置是否与已放置的皇后冲突
			if cols[col] || diag1[row-col+n-1] || diag2[row+col] {
				continue
			}

			// 放置皇后
			queens[row] = col
			cols[col] = true
			diag1[row-col+n-1] = true
			diag2[row+col] = true

			// 递归处理下一行
			backtrack(row + 1)

			// 回溯，撤销当前行的皇后放置
			cols[col] = false
			diag1[row-col+n-1] = false
			diag2[row+col] = false
		}
	}

	// 从第0行开始回溯
	backtrack(0)
	return result
}

// generateBoard 将皇后位置转换为字符串表示的棋盘
func generateBoard(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if queens[i] == j {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board[i] = string(row)
	}
	return board
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=4",
			n:        4,
			expected: 2,
		},
		{
			name:     "n=8",
			n:        8,
			expected: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solveNQueens(tt.n)

			if len(result) != tt.expected {
				t.Errorf("solveNQueens(%d) = %d solutions, expected %d",
					tt.n, len(result), tt.expected)
				return
			}

			// 验证解的正确性
			for i, solution := range result {
				if !isValidSolution(solution, tt.n) {
					t.Errorf("Solution %d is invalid: %v", i, solution)
				}
			}
		})
	}

	// 特殊测试：n=2和n=3应该无解
	noSolutionTests := []int{2, 3}
	for _, n := range noSolutionTests {
		result := solveNQueens(n)
		if len(result) != 0 {
			t.Errorf("solveNQueens(%d) should have no solution, but got %d", n, len(result))
		}
	}
}

// isValidSolution 验证解的正确性
func isValidSolution(board []string, n int) bool {
	// 检查棋盘大小
	if len(board) != n {
		return false
	}

	queens := make([][2]int, 0)
	for i := 0; i < n; i++ {
		if len(board[i]) != n {
			return false
		}

		for j := 0; j < n; j++ {
			if board[i][j] == 'Q' {
				queens = append(queens, [2]int{i, j})
			}
		}
	}

	// 检查皇后数量
	if len(queens) != n {
		return false
	}

	// 检查是否互相攻击
	for i := 0; i < len(queens); i++ {
		for j := i + 1; j < len(queens); j++ {
			r1, c1 := queens[i][0], queens[i][1]
			r2, c2 := queens[j][0], queens[j][1]

			// 检查行、列、对角线冲突
			if r1 == r2 || c1 == c2 || abs(r1-r2) == abs(c1-c2) {
				return false
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
