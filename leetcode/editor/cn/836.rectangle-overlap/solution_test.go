package leetcode

import "testing"

// 836.矩形重叠

/**
矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。矩形的上下边平行于 x 轴，
左右边平行于 y 轴。

 如果相交的面积为 正 ，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。

 给出两个矩形 rec1 和 rec2 。如果它们重叠，返回 true；否则，返回 false 。



 示例 1：


输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]
输出：true


 示例 2：


输入：rec1 = [0,0,1,1], rec2 = [1,0,2,1]
输出：false


 示例 3：


输入：rec1 = [0,0,1,1], rec2 = [2,2,3,3]
输出：false




 提示：


 rect1.length == 4
 rect2.length == 4
 -10⁹ <= rec1[i], rec2[i] <= 10⁹
 rec1 和 rec2 表示一个面积不为零的有效矩形


 Related Topics 几何 数学 👍 321 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 检查x轴方向是否重叠：rec1的右边界大于rec2的左边界 且 rec2的右边界大于rec1的左边界
	xOverlap := rec1[2] > rec2[0] && rec2[2] > rec1[0]

	// 检查y轴方向是否重叠：rec1的上边界大于rec2的下边界 且 rec2的上边界大于rec1的下边界
	yOverlap := rec1[3] > rec2[1] && rec2[3] > rec1[1]

	// 只有当x和y方向都重叠时，矩形才重叠
	return xOverlap && yOverlap
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRectangleOverlap(t *testing.T) {
	tests := []struct {
		rec1     []int
		rec2     []int
		expected bool
	}{
		{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
		{[]int{7, 8, 13, 15}, []int{10, 8, 12, 20}, true},
	}

	for i, test := range tests {
		result := isRectangleOverlap(test.rec1, test.rec2)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, test.expected, result)
		}
	}
}
