package leetcode

// 228.汇总区间

import (
	"fmt"
	"testing"
)

/**
给定一个 无重复元素 的 有序 整数数组 nums 。

 区间 [a,b] 是从 a 到 b（包含）的所有整数的集合。

 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个区间但不属于
nums 的数字 x 。

 列表中的每个区间范围 [a,b] 应该按如下格式输出：


 "a->b" ，如果 a != b
 "a" ，如果 a == b




 示例 1：


输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"


 示例 2：


输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"




 提示：


 0 <= nums.length <= 20
 -2³¹ <= nums[i] <= 2³¹ - 1
 nums 中的所有值都 互不相同
 nums 按升序排列


 Related Topics 数组 👍 462 👎 0

*/
//leetcode submit region begin(Prohibit modification and deletion)
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start := nums[0]
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		// 如果当前数字与前一个数字连续
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			// 不连续，保存当前区间并开始新区间
			if start == end {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
			end = nums[i]
		}
	}

	// 处理最后一个区间
	if start == end {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
		{[]int{1, 3, 5, 7}, []string{"1", "3", "5", "7"}},
		{[]int{-2, -1, 0, 1, 3}, []string{"-2->1", "3"}},
	}

	for i, test := range tests {
		result := summaryRanges(test.nums)
		if !stringSlicesEqual(result, test.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v. Input: %v",
				i+1, test.expected, result, test.nums)
		}
	}
}

// 辅助函数：比较两个字符串切片是否相等
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
