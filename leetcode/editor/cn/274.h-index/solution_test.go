package leetcode

// 274.H 指数

import (
	"testing"
)

/**
给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大
于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。



 示例 1：


输入：citations = [3,0,6,1,5]
输出：3
解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。

 示例 2：


输入：citations = [1,3,1]
输出：1




 提示：


 n == citations.length
 1 <= n <= 5000
 0 <= citations[i] <= 1000


 Related Topics 数组 计数排序 排序 👍 612 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func hIndex(citations []int) int {
	n := len(citations)
	// 创建桶数组，索引表示引用次数，值表示该引用次数的论文数量
	// 由于引用次数最大为1000，且最多有n篇论文，所以桶大小为n+1即可
	buckets := make([]int, n+1)

	// 统计每个引用次数的论文数量
	for i := 0; i < n; i++ {
		if citations[i] >= n {
			// 引用次数超过n的都放在buckets[n]中
			buckets[n]++
		} else {
			buckets[citations[i]]++
		}
	}
	// 从后往前累加，计算引用次数大于等于i的论文总数
	count := 0
	for i := n; i >= 0; i-- {
		// 如果引用次数大于等于i的论文数量不少于i篇，则h指数为i
		count += buckets[i]
		if count >= i {
			return i
		}
	}

	return 0

}

//leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		expected  int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{100}, 1},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 1, 1, 1}, 1},
	}

	for i, test := range tests {
		result := hIndex(test.citations)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
