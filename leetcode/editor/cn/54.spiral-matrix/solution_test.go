package leetcode

// 54.螺旋矩阵

import (
	"fmt"
	"reflect"
	"testing"
)

/**
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



 示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]


 示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]




 提示：


 m == matrix.length
 n == matrix[i].length
 1 <= m, n <= 10
 -100 <= matrix[i][j] <= 100


 Related Topics 数组 矩阵 模拟 👍 1976 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	left := 0
	right := n - 1
	bottom := m - 1
	top := 0
	k := 0

	for {
		for i := left; i <= right; i++ {
			ans[k] = matrix[top][i]
			k++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			ans[k] = matrix[i][right]
			k++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			ans[k] = matrix[bottom][i]
			k++
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			ans[k] = matrix[i][left]
			k++
		}
		left++
		if left > right {
			break
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestSpiralMatrix(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		matrix [][]int
		ans    []int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, test := range tests {
		ans := spiralOrder(test.matrix)
		if !reflect.DeepEqual(ans, test.ans) {
			t.Errorf("matrix: %v, ans: %v, actual: %v", test.matrix, test.ans, ans)
		}
	}
}
