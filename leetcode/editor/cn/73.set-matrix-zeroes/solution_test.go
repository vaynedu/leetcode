package leetcode

// 73.矩阵置零

import (
	"fmt"
	"testing"
)

/**
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。






 示例 1：


输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]


 示例 2：


输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]




 提示：


 m == matrix.length
 n == matrix[0].length
 1 <= m, n <= 200
 -2³¹ <= matrix[i][j] <= 2³¹ - 1




 进阶：


 一个直观的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 你能想出一个仅使用常量空间的解决方案吗？


 Related Topics 数组 哈希表 矩阵 👍 1258 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 遍历数组，找到0 并且记录要置0的 行列
	rowZeroMap := make(map[int]bool)
	colZeroMap := make(map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowZeroMap[i] = true
				colZeroMap[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowZeroMap[i] || colZeroMap[j] {
				matrix[i][j] = 0
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSetMatrixZeroes(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "test1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "test2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.matrix)
			setZeroes(tt.matrix)
			t.Log(tt.matrix)
		})
	}
}
