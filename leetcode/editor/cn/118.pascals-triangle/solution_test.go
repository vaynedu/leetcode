package leetcode

// 118.杨辉三角

import (
	"testing"
)

/**
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

 在「杨辉三角」中，每个数是它左上方和右上方的数的和。





 示例 1:


输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]


 示例 2:


输入: numRows = 1
输出: [[1]]




 提示:


 1 <= numRows <= 30


 👍 1290 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	// 创建结果切片
	triangle := make([][]int, numRows)

	// 逐行构造杨辉三角
	for i := 0; i < numRows; i++ {
		// 每行有i+1个元素
		row := make([]int, i+1)

		// 每行的第一个和最后一个元素都是1
		row[0] = 1
		row[i] = 1

		// 计算中间的元素值
		for j := 1; j < i; j++ {
			// 当前元素 = 上一行左上方元素 + 上一行右上方元素
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		// 将当前行添加到结果中
		triangle[i] = row
	}

	return triangle
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPascalsTriangle(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:    "示例1: numRows=5",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "示例2: numRows=1",
			numRows: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "numRows=0",
			numRows:  0,
			expected: [][]int{},
		},
		{
			name:    "numRows=2",
			numRows: 2,
			expected: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "numRows=6",
			numRows: 6,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generate(tt.numRows)

			if !equalTriangle(result, tt.expected) {
				t.Errorf("generate(%d) = %v, expected %v", tt.numRows, result, tt.expected)
			}
		})
	}
}

// equalTriangle 比较两个杨辉三角是否相等
func equalTriangle(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
