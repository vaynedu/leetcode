# 动态规划题目

| 标题                                                                     | 内部跳转                                                                                   | 思路   |
|------------------------------------------------------------------------|----------------------------------------------------------------------------------------|------|
| [416.分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/) | [416.分割等和子集](../../leetcode/editor/cn/416.partition-equal-subset-sum/solution_test.go) | 01背包 | 01背包 dp[j] = max(dp[j], dp[j - nums[i]] + nums[i]) |
