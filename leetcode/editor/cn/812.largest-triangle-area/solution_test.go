package leetcode

// 812.最大三角形面积

import (
	"testing"
)

/**
给你一个由 X-Y 平面上的点组成的数组 points ，其中 points[i] = [xi, yi] 。从其中取任意三个不同的点组成三角形，返回能组成的最大
三角形的面积。与真实值误差在 10⁻⁵ 内的答案将会视为正确答案。



 示例 1：


输入：points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
输出：2.00000
解释：输入中的 5 个点如上图所示，红色的三角形面积最大。


 示例 2：


输入：points = [[1,0],[0,0],[0,1]]
输出：0.50000




 提示：


 3 <= points.length <= 50
 -50 <= xi, yi <= 50
 给出的所有点 互不相同


 Related Topics 几何 数组 数学 👍 220 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)

	// 遍历所有可能的三个点的组合
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// 计算三个点组成的三角形面积
				area := calculateTriangleArea(points[i], points[j], points[k])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// 使用向量叉积计算三角形面积
// 公式: 面积 = 0.5 * |x1(y2-y3) + x2(y3-y1) + x3(y1-y2)|
func calculateTriangleArea(p1, p2, p3 []int) float64 {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]
	x3, y3 := p3[0], p3[1]

	area := 0.5 * float64(abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
	return area
}

// 计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLargestTriangleArea(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected float64
	}{
		{
			name:     "示例1",
			points:   [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			expected: 2.0,
		},
		{
			name:     "示例2",
			points:   [][]int{{1, 0}, {0, 0}, {0, 1}},
			expected: 0.5,
		},
		{
			name:     "三点共线",
			points:   [][]int{{0, 0}, {1, 1}, {2, 2}},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestTriangleArea(tt.points)
			if result != tt.expected {
				t.Errorf("largestTriangleArea(%v) = %f; expected %f", tt.points, result, tt.expected)
			}
		})
	}
}
