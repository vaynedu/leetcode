package leetcode

// 56.合并区间

import (
	"sort"
	"testing"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回
一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



 示例 1：


输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].


 示例 2：


输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

 示例 3：


输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。




 提示：


 1 <= intervals.length <= 10⁴
 intervals[i].length == 2
 0 <= starti <= endi <= 10⁴


 Related Topics 数组 排序 👍 2667 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res

}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{[][]int{{1, 3}}, [][]int{{1, 3}}},
		{[][]int{}, [][]int{}},
	}

	for i, test := range tests {
		result := merge(test.intervals)
		if !intervalSlicesEqual(result, test.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v. Input: %v",
				i+1, test.expected, result, test.intervals)
		}
	}
}

// 辅助函数：比较两个区间切片是否相等
func intervalSlicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) || a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
