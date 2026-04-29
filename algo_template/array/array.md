# 数组题目

数组是最常见的载体，本身不是单一算法模式。看到数组题时，先判断它落在哪个模式：哈希、排序、双指针、滑动窗口、二分、前缀和、DP。

## 识别信号

- 查找两个数/重复元素：优先 hash。
- 已排序或可排序后处理：优先二分、双指针、贪心。
- 连续子数组/子串：优先滑动窗口、前缀和、DP。
- 原地移动或分区：优先双指针。

## 核心不变量

- hash：已经遍历过的元素可以 O(1) 查询。
- 排序+双指针：被跳过的候选区间不会产生答案。
- 前缀和：区间和可由两个前缀差得到。
- 原地分区：每个区间维护明确含义，例如小于区、等于区、大于区。

## 决策流程

```text
1. 先问是否需要连续区间。
2. 再问是否可以排序。
3. 再问是否需要 O(1) 查询历史元素。
4. 最后根据目标选择 hash、双指针、二分、窗口、DP。
```

以下是一些 LeetCode 上经典且出现频率较高的数组相关题目：

| 标题                                                                                                            | 频率 | 思路                     |
|---------------------------------------------------------------------------------------------------------------|----|------------------------|
| [1.两数之和](https://leetcode.cn/problems/two-sum/description/)                   | 高  | 哈希表，存储已遍历元素及其索引              |
| [15.三数之和](https://leetcode.cn/problems/3sum/description/)                   | 高  | 排序 + 双指针，固定一个数，找另外两个数之和为其相反数              |
| [217.存在重复元素](https://leetcode.cn/problems/contains-duplicate/description/)                   | 高  | 哈希表，检查元素是否已存在              |
| [53.最大子数组和](https://leetcode.cn/problems/maximum-subarray/description/)                   | 高  | 动态规划，记录以当前元素结尾的最大子数组和              |
| [238.除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/description/)                   | 高  | 左右乘积列表，分别计算元素左侧和右侧的乘积              |
| [75.颜色分类](https://leetcode.cn/problems/sort-colors/description/)                   | 高  | 双指针，荷兰国旗问题，将数组分为三个部分              |
| [33.搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/description/)                   | 高  | 二分查找，判断旋转点和目标值的位置              |
| [121.买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/)                   | 高  | 一次遍历，记录最小买入价格和最大利润              |
| [48.旋转图像](https://leetcode.cn/problems/rotate-image/description/)                   | 高  | 先转置再翻转，或分块旋转              |
| [11.盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/)                   | 高  | 双指针，向中间移动较短的指针以寻找更大的面积              | 
