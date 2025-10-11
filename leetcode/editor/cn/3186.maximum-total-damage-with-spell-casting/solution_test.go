package leetcode

// 3186.施咒的最大总伤害

import (
	"sort"
	"testing"
)

/**
一个魔法师有许多不同的咒语。

 给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。

 已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1
 或者 power[i] + 2 的咒语。

 每个咒语最多只能被使用 一次 。

 请你返回这个魔法师可以达到的伤害值之和的 最大值 。



 示例 1：


 输入：power = [1,1,3,4]


 输出：6

 解释：

 可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。

 示例 2：


 输入：power = [7,1,6,6]


 输出：13

 解释：

 可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。



 提示：


 1 <= power.length <= 10⁵
 1 <= power[i] <= 10⁹


 Related Topics 数组 哈希表 双指针 二分查找 动态规划 计数 排序 👍 81 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*
这个问题可以通过以下步骤解决：
预处理：统计每种伤害值的总伤害量
排序：按伤害值大小排序
动态规划：对每个伤害值决定是否选择，在满足约束条件下求最大值*/
func maximumTotalDamage(power []int) int64 {
	// 统计每个伤害值出现的次数
	count := make(map[int]int)
	for _, p := range power {
		count[p]++
	}

	// 获取所有不同的伤害值并排序
	uniquePowers := make([]int, 0, len(count))
	for p := range count {
		uniquePowers = append(uniquePowers, p)
	}
	sort.Ints(uniquePowers)

	n := len(uniquePowers)
	// dp[i] 表示考虑前i个不同伤害值能获得的最大伤害
	dp := make([]int64, n+1)

	for i := 1; i <= n; i++ {
		currentPower := uniquePowers[i-1]
		currentTotalDamage := int64(currentPower * count[currentPower])

		// 不选择当前伤害值
		dp[i] = dp[i-1]

		// 选择当前伤害值，需要找到可以与之共存的前一个状态
		// 即找到最大的j使得uniquePowers[j-1] <= currentPower - 3
		// 因为如果选择当前伤害值currentPower，就不能选择范围[currentPower-2, currentPower+2]内的伤害值
		j := i - 1
		for j >= 1 && uniquePowers[j-1] >= currentPower-2 {
			j--
		}

		// 选择当前伤害值的最大收益
		dp[i] = max(dp[i], dp[j]+currentTotalDamage)
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumTotalDamageWithSpellCasting(t *testing.T) {
	tests := []struct {
		power    []int
		expected int64
	}{
		{[]int{1, 1, 3, 4}, 6},
		{[]int{7, 1, 6, 6}, 13},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{5, 7, 8, 9, 10}, 15}, // 选择5和10
	}

	for i, test := range tests {
		result := maximumTotalDamage(test.power)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		} else {
			t.Logf("Test case %d passed: input=%v, output=%d", i+1, test.power, result)
		}
	}
}
