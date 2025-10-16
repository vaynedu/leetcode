package leetcode

// 167.两数之和 II - 输入有序数组

import (
	"testing"
)

/**
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个
数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.
length 。

 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

 你所设计的解决方案必须只使用常量级的额外空间。

 示例 1：


输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

 示例 2：


输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

 示例 3：


输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。




 提示：


 2 <= numbers.length <= 3 * 10⁴
 -1000 <= numbers[i] <= 1000
 numbers 按 非递减顺序 排列
 -1000 <= target <= 1000
 仅存在一个有效答案


 Related Topics 数组 双指针 二分查找 👍 1346 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	// 思路
	// 只要两个数的下标， 非递减的顺序

	// map key 值， value 存储下标，  要看下是否有重复的数字， 这里要重新处理key

	m := make(map[int][]int)

	for i := 0; i < len(numbers); i++ {
		if _, exist := m[numbers[i]]; !exist {
			m[numbers[i]] = make([]int, 0)
		}
		m[numbers[i]] = append(m[numbers[i]], i) // 存储下标
	}

	for i := 0; i < len(numbers); i++ {
		values, ok := m[target-numbers[i]]
		if !ok {
			continue
		}

		if values[0] == i && len(values) > 1 {
			return []int{i + 1, values[1] + 1}
		}

		return []int{i + 1, values[0] + 1}

	}

	return []int{-1, -1}
}

func twoSumV1(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right++
		}
	}
	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSumIiInputArrayIsSorted(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 8, []int{3, 5}},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, 0, []int{5, 6}},
	}

	for _, test := range tests {
		result := twoSum(test.numbers, test.target)
		if len(result) != 2 || result[0] != test.expected[0] || result[1] != test.expected[1] {
			t.Errorf("Input: numbers=%v, target=%d, Expected: %v, Got: %v",
				test.numbers, test.target, test.expected, result)
		}
	}
}
