package leetcode

// 128.最长连续序列

import (
	"testing"
)

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



 示例 1：


输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

 示例 2：


输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


 示例 3：


输入：nums = [1,0,1,2]
输出：3




 提示：


 0 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 并查集 数组 哈希表 👍 2664 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

/*
要实现O(n)时间复杂度，我们需要使用哈希表来优化查找过程：
1.将所有数字放入集合中，实现O(1)的查找
2.对于每个数字，判断它是否为某个连续序列的起点（即num-1不在集合中）
3.如果是起点，则向后查找连续数字，统计序列长度
4.记录并更新最大长度
*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// 向后查找连续数字
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// 更新最大长度
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}

	for i, test := range tests {
		result := longestConsecutive(test.nums)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d. Input: %v",
				i+1, test.expected, result, test.nums)
		}
	}
}
