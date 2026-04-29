# 数组 · 题目索引

> 数组不是单一模式。这个索引用来快速判断数组题应转向哪个模式。

## 模板要点

```text
查找/计数：hash
有序或可排序：二分、双指针、贪心
连续区间：滑动窗口、前缀和、DP
原地分区：双指针、荷兰国旗
区间合并：排序后维护当前合并段
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 15 | [三数之和](./../../interview/15.3sum.go) | 中等 | ★★★ | 排序+双指针 | `go test -run TestThreeSum` |
| 33 | [搜索旋转排序数组](./../../interview/33.search-in-rotated-sorted-array.go) | 中等 | ★★★ | 二分 | `go test -run TestSearch33` |
| 53 | [最大子数组和](./../../interview/53.maximum-subarray.go) | 中等 | ★★★ | DP/Kadane | `go test -run TestMaxSubArray` |
| 56 | [合并区间](./../../interview/56.merge-intervals.go) | 中等 | ★★★ | 排序+区间 | 待补测试 |
| 238 | [除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) | 中等 | ★★★ | 前后缀乘积 | 待补 |
| 75 | [颜色分类](https://leetcode.cn/problems/sort-colors/) | 中等 | ★★☆ | 原地分区 | 待补 |
| 1 | [两数之和](https://leetcode.cn/problems/two-sum/) | 简单 | ★★★ | Hash | 待补 |

## 学习路径建议

```text
第一步：1 / 15（hash 与排序双指针）
  ↓
第二步：33（有序数组的二分判断）
  ↓
第三步：53（连续子数组 DP）
  ↓
第四步：56 / 75 / 238（区间、分区、前后缀）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'Test(ThreeSum|Search33|MaxSubArray)' -count=1
```
