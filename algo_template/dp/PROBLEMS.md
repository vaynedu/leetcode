# 动态规划 · 题目索引

> 学完 `dp.go` 模板后，用这些题实战检验。

## 模板要点

```
DP 四步：
1. 定义状态  dp[i] / dp[i][j]
2. 转移方程  dp[i] = f(dp[i-1], ...)
3. 初始化    dp[0] = ...
4. 遍历顺序  从小到大

经典模型：
  - 斐波那契：dp[i]=dp[i-1]+dp[i-2]
  - 路径DP：dp[i][j]=dp[i-1][j]+dp[i][j-1]
  - 爬楼梯：dp[i]=dp[i-1]+dp[i-2]
  - 卡特兰数：dp[n]=Σdp[i]*dp[n-1-i]
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 53 | [最大子数组和](./../../interview/53.maximum-subarray.go) | 简单 | ★★★ | Kadane/一维 | `go test -run TestMaxSubarray` |
| 62 | [不同路径](./../../interview/62.unique-paths.go) | 中等 | ★★★ | 路径DP | `go test -run TestUniquePaths` |
| 63 | [不同路径 II](./../../interview/63.unique-paths-ii.go) | 中等 | ★★☆ | 路径DP+障碍 | `go test -run TestUniquePathsWithObstacles` |
| 64 | [最小路径和](./../../interview/64.minimum-path-sum.go) | 中等 | ★★★ | 路径DP+最小 | `go test -run TestMinPathSum` |
| 70 | [爬楼梯](./../../interview/70.climbing-stairs.go) | 简单 | ★★★ | 斐波那契 | `go test -run TestClimbStairs` |
| 96 | [不同的二叉搜索树](./../../interview/96.unique-binary-search-trees.go) | 中等 | ★★★ | 卡特兰数 | `go test -run TestNumTrees` |
| 120 | [三角形最小路径和](./../../interview/120.triangle.go) | 中等 | ★★★ | 路径DP+上行 | `go test -run TestMinimumTotal` |

## 学习路径建议

```
第一步：70 / 53（斐波那契/Kadane，最简单的DP）
  ↓
第二步：62 / 63 / 64（路径DP系列）
  ↓
第三步：96（卡特兰数，有一定难度）
  ↓
第四步：120（从底往上行优化）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/dp/ -v
```
