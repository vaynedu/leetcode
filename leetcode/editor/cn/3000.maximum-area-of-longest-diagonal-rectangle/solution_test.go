package leetcode

// 3000.对角线最长的矩形的面积

import (
	"fmt"
	"testing"
)

/**
给你一个下标从 0 开始的二维整数数组 dimensions。

 对于所有下标 i（0 <= i < dimensions.length），dimensions[i][0] 表示矩形 i 的长度，而 dimensions[
i][1] 表示矩形 i 的宽度。

 返回对角线最 长 的矩形的 面积 。如果存在多个对角线长度相同的矩形，返回面积最 大 的矩形的面积。



 示例 1：


输入：dimensions = [[9,3],[8,6]]
输出：48
解释：
下标 = 0，长度 = 9，宽度 = 3。对角线长度 = sqrt(9 * 9 + 3 * 3) = sqrt(90) ≈
  9.487。
下标 = 1，长度 = 8，宽度 = 6。对角线长度 = sqrt(8 * 8 + 6 * 6) = sqrt(100) = 10。
因此，下标为 1 的矩形对角线更长，所以返回面积 = 8 * 6 = 48。


 示例 2：


输入：dimensions = [[3,4],[4,3]]
输出：12
解释：两个矩形的对角线长度相同，为 5，所以最大面积 = 12。




 提示：


 1 <= dimensions.length <= 100
 dimensions[i].length == 2
 1 <= dimensions[i][0], dimensions[i][1] <= 100


 Related Topics 数组 👍 18 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func areaOfMaxDiagonal(dimensions [][]int) int {
	ans := 0  // 最大面积
	sqrt := 0 // 最长对角线的平方
	for _, dimension := range dimensions {
		// 计算当前矩形对角线长度的平方
		curSqrt := dimension[0]*dimension[0] + dimension[1]*dimension[1]

		if curSqrt < sqrt {
			// 当前对角线更短，跳过
			continue
		} else if curSqrt > sqrt {
			sqrt = curSqrt
			// 当前对角线更长，更新面积
			ans = dimension[0] * dimension[1]
		} else {
			// 对角线相等，取面积较大者
			ans = max(ans, dimension[0]*dimension[1])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumAreaOfLongestDiagonalRectangle(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name: "test1",
			input: [][]int{
				{2, 6},
				{5, 1},
				{3, 10},
				{8, 4},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := areaOfMaxDiagonal(tt.input)
			if got != tt.want {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
