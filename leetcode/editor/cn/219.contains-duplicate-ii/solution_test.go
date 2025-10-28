package leetcode

// 219.存在重复元素 II

import (
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i -
j) <= k 。如果存在，返回 true ；否则，返回 false 。



 示例 1：


输入：nums = [1,2,3,1], k = 3
输出：true

 示例 2：


输入：nums = [1,0,1,1], k = 1
输出：true

 示例 3：


输入：nums = [1,2,3,1,2,3], k = 2
输出：false





 提示：


 1 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 0 <= k <= 10⁵


 Related Topics 数组 哈希表 滑动窗口 👍 825 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	// 某个窗口内，是否有两个相等值
	// 0,3   3,

	// 使用哈希表存储每个元素最近出现的索引
	indexMap := make(map[int]int)
	for i, num := range nums {
		if prevIndex, exist := indexMap[num]; exist {
			if i-prevIndex <= k {
				return true
			}
		}
		indexMap[num] = i
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainsDuplicateIi(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{99, 99}, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
	}

	for i, test := range tests {
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %v, got %v. Input: nums=%v, k=%d",
				i+1, test.expected, result, test.nums, test.k)
		}
	}
}
