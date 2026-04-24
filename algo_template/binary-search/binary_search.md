# 二分查找（Binary Search）

## 一、核心思想

二分查找通过**每次将搜索范围缩小一半**来快速定位目标，时间复杂度 O(log N)。

```
前提：数组必须有序

      L           M           R
      ↓           ↓           ↓
    [1, 3, 5, 7, 9, 11, 13]
    
    target = 7
    mid = 5 < 7 → L = mid + 1
         L   M           R
         ↓   ↓           ↓
    [1, 3, 5, 7, 9, 11, 13]
    
    mid = 9 > 7 → R = mid - 1
         L   M   R
         ↓   ↓   ↓
    [1, 3, 5, 7, 9, 11, 13]
    
    mid = 7 == 7 → 找到！
```

---

## 二、三种模板

### 模板1: 标准二分（左闭右闭 [L, R]）

**适用场景**：查找**存在且唯一**的元素

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1  // ★ 左闭右闭
    for left <= right {            // ★ <=
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

**代表题目**：704. 二分查找

---

### 模板2: 左侧边界（左闭右开 [L, R)）

**适用场景**：找第一个 >= target 的位置

```go
func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)   // ★ 右开
    for left < right {             // ★ <
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid            // ★ 不减1
        }
    }
    return left                    // ★ 返回 left
}
```

**代表题目**：
- 35. 搜索插入位置
- 278. 第一个错误的版本

---

### 模板3: 右侧边界（左闭右开 [L, R)）

**适用场景**：找最后一个 <= target 的位置

```go
func upperBound(nums []int, target int) int {
    left, right := 0, len(nums)   // 右开
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1        // ★ <= 时 left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
```

---

## 三、模板选择决策树

```
二分查找问题
    │
    ├─ 找精确值（元素存在且唯一）
    │   └─ 模板1：left <= right, left = mid+1, right = mid-1
    │
    ├─ 找左边界（第一个 >= target）
    │   └─ 模板2：left < right, left = mid+1, right = mid
    │
    ├─ 找右边界（最后一个 <= target）
    │   └─ 模板3：left < right, left = mid+1, right = mid
    │
    └─ 旋转数组
        └─ 先判断哪半边有序，再决定搜索范围
```

---

## 四、通用口诀

```
二分查找三要素
左闭右闭或右开
mid 防溢要记牢
等不等号决定边界
缩小范围不断收敛
直至找到目标值
或返回 -1 表不存在
```

---

## 五、易错点汇总

| 题目 | 易错点 |
|------|--------|
| 704 二分查找 | while 条件是 `<=` 不是 `<` |
| 35 搜索插入位置 | 返回 left（第一个 >= target）|
| 69 平方根 | mid*mid 会溢出，用 `mid <= x/mid` |
| 33 搜索旋转数组 | 先判断哪半边有序 |
| 162 寻找峰值 | nums[mid] < nums[mid+1] 时往右走 |

---

## 六、题目速查表

| 难度 | 题号 | 名称 | 模板 |
|------|------|------|------|
| 简单 | 704 | 二分查找 | 模板1 |
| 简单 | 35 | 搜索插入位置 | 模板2 |
| 简单 | 278 | 第一个错误的版本 | 模板2 |
| 简单 | 69 | x 的平方根 | 边界 |
| 简单 | 367 | 有效的完全平方数 | 边界 |
| 中等 | 33 | 搜索旋转排序数组 | 旋转 |
| 中等 | 153 | 寻找旋转排序数组中的最小值 | 旋转 |
| 中等 | 162 | 寻找峰值 | 峰值 |

---

## 七、复杂度总结

| 类型 | 时间 | 空间 |
|------|------|------|
| 标准二分 | O(log N) | O(1) |
| 边界二分 | O(log N) | O(1) |
| 旋转数组 | O(log N) | O(1) |
