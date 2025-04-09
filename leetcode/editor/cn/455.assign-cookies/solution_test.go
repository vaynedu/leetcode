package leetcode

// 455.分发饼干

import (
	"sort"
	"testing"
)

/**
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，
我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是满足尽可能多的孩子，并输出这个最大数值。

 示例 1:


输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3 个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是 1，你只能让胃口值是 1 的孩子满足。
所以你应该输出 1。


 示例 2:


输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2 个孩子的胃口值分别是 1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出 2。




 提示：


 1 <= g.length <= 3 * 10⁴
 0 <= s.length <= 3 * 10⁴
 1 <= g[i], s[j] <= 2³¹ - 1




 注意：本题与 2410. 运动员和训练师的最大匹配数 题相同。

 Related Topics 贪心 数组 双指针 排序 👍 923 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	ans := 0
	sort.Ints(g)
	sort.Ints(s)

	for i, j := len(s)-1, len(g)-1; i >= 0 && j >= 0; {
		if s[i] >= g[j] {
			ans++
			j--
			i--
		} else {
			j--
		}

	}
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestAssignCookies(t *testing.T) {
	tests := []struct {
		g        []int
		s        []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{}, []int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, []int{}, 0},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 3},
		{[]int{10, 20, 30}, []int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, []int{10, 20, 30}, 3},
	}

	for _, tt := range tests {
		result := findContentChildren(tt.g, tt.s)
		if result != tt.expected {
			t.Errorf("findContentChildren(%v, %v) = %d; expected %d", tt.g, tt.s, result, tt.expected)
		}
	}
}
