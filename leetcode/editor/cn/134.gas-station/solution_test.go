package leetcode

// 134.加油站

import (
	"testing"
)

/**
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。




 示例 1:


输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

 示例 2:


输入: gas = [2,3,4], cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。



 提示:


 n == gas.length == cost.length
 1 <= n <= 10⁵
 0 <= gas[i], cost[i] <= 10⁴
 输入保证答案唯一。


 Related Topics 贪心 数组 👍 1902 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

// 解题思路
// 加油站问题可以通过贪心算法来解决。核心思想是：
// 如果总油量小于总消耗量，那么肯定无法完成一圈
// 如果可以从某个站点出发，那么从该站点到任一站点的累积油量都不会为负
// 使用贪心策略：一旦发现从某起点开始油量变为负数，说明该起点及之前的所有点都不能作为起点，应该从下一个点重新开始
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currentGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start

}

//leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestGasStation(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, 0},
	}

	for i, test := range tests {
		result := canCompleteCircuit(test.gas, test.cost)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
