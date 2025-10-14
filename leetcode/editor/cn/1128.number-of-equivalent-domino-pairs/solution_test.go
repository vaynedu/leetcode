package leetcode

// 1128.等价多米诺骨牌对的数量

import (
	"testing"
)

/**
给你一组多米诺骨牌 dominoes 。

 形式上，dominoes[i] = [a, b] 与 dominoes[j] = [c, d] 等价 当且仅当 (a == c 且 b == d) 或者 (
a == d 且 b == c) 。即一张骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌。

 在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i,
j) 的数量。



 示例 1：


输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
输出：1


 示例 2：


输入：dominoes = [[1,2],[1,2],[1,1],[1,2],[2,2]]
输出：3




 提示：


 1 <= dominoes.length <= 4 * 10⁴
 dominoes[i].length == 2
 1 <= dominoes[i][j] <= 9


 Related Topics 数组 哈希表 计数 👍 201 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numEquivDominoPairs(dominoes [][]int) int {
	if len(dominoes) == 1 {
		return 0
	}

	countMap := make(map[[2]int]int)

	for _, dominoe := range dominoes {
		var key [2]int
		if dominoe[0] <= dominoe[1] {
			key[0], key[1] = dominoe[0], dominoe[1]
		} else {
			key[0], key[1] = dominoe[1], dominoe[0]
		}
		countMap[key]++
	}
	// 计算等价对的数量
	result := 0
	for _, count := range countMap {
		if count > 1 {
			result += count * (count - 1) / 2
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfEquivalentDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}, 3},
		{[][]int{{1, 2}}, 0},
		{[][]int{{1, 2}, {2, 1}}, 1},
	}

	for i, test := range tests {
		actual := numEquivDominoPairs(test.dominoes)
		if actual != test.expected {
			t.Errorf("Test %d failed: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}
