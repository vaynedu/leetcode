# 贪心 · 题目索引

> 学完 `greedy.md` 后，用这些题训练“为什么局部选择可以交换成全局最优”。

## 模板要点

```text
交换论证：最优解没选当前选择时，交换后不变差
区间贪心：按右端点、左端点或高度排序，减少未来约束
前缀贪心：维护当前最优前缀或最低成本
可达性贪心：维护当前能到达的最远位置
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 53 | [最大子数组和](./../../interview/53.maximum-subarray.go) | 中等 | ★★★ | Kadane/丢弃负前缀 | `go test -run TestMaxSubArray` |
| 56 | [合并区间](./../../interview/56.merge-intervals.go) | 中等 | ★★★ | 排序+区间合并 | 待补测试 |
| 121 | [买卖股票的最佳时机](./../../doc/algorithm-visualizations/121.best-time-to-buy-and-sell-stock.html) | 简单 | ★★★ | 最低买入价 | 可视化 |
| 122 | [买卖股票的最佳时机 II](./../../doc/algorithm-visualizations/122.best-time-to-buy-and-sell-stock-ii.html) | 中等 | ★★☆ | 累加正收益 | 可视化 |
| 435 | [无重叠区间](./../../doc/algorithm-visualizations/435.non-overlapping-intervals.html) | 中等 | ★★☆ | 最小右端点 | 可视化 |
| 455 | [分发饼干](./../../doc/algorithm-visualizations/455.assign-cookies.html) | 简单 | ★★☆ | 小胃口配小饼干 | 可视化 |
| 763 | [划分字母区间](./../../doc/algorithm-visualizations/763.partition-labels.html) | 中等 | ★★☆ | 最远出现位置 | 可视化 |

## 学习路径建议

```text
第一步：53 / 121（前缀最优）
  ↓
第二步：455（排序后局部匹配）
  ↓
第三步：56 / 435（区间贪心）
  ↓
第四步：763（维护当前段最远边界）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'TestMaxSubArray' -count=1
```
