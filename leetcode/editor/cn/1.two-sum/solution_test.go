package leetcode

// 两数之和

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
//
//
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
//
// Related Topics 数组 哈希表 👍 19075 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	expected := []int{0, 1}
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	expected2 := []int{1, 2}
	if result2[0] != expected2[0] || result2[1] != expected2[1] {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}

	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	expected3 := []int{0, 1}
	if result3[0] != expected3[0] || result3[1] != expected3[1] {
		t.Errorf("Expected %v, got %v", expected3, result3)
	}

	fmt.Println("All tests passed!")
}
