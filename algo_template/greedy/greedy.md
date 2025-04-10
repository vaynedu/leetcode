# 贪心算法题目
定义：贪心算法是一种在每一步选择中都选择在当前状态下最好或最优的选项的算法。

| 标题                                                           | 内部跳转                                                                         | 思路                                                                |
|--------------------------------------------------------------|------------------------------------------------------------------------------|-------------------------------------------------------------------|
| [455.分发饼干](https://leetcode.cn/problems/assign-cookies/)     | [455.分发饼干](../../leetcode/editor/cn/455.assign-cookies/solution_test.go)     | 同时遍历两个数组，按照贪心算法，最大饼干喂给可以满足的孩子                                     |
| [376.摆动序列](https://leetcode.cn/problems/wiggle-subsequence/) | [376.摆动序列](../../leetcode/editor/cn/376.wiggle-subsequence/solution_test.go) | 贪心算法，波峰和波谷的差值，如果差值大于0，则当前位置为波峰，否则为波谷                              |
| [53.最大子数组和](https://leetcode.cn/problems/maximum-subarray/)  | [53.最大子数组和](../../leetcode/editor/cn/53.maximum-subarray/solution_test.go)   | 贪心算法，如果当前位置的值加上前一个位置的值大于前一个位置的值，则当前位置的值加上前一个位置的值，否则当前位置的值为前一个位置的值 |