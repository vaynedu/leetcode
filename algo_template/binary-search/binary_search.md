# 二分查找速查手册

## 一、选择决策树

```
看到题目
  │
  ├─ 有序数组找 target ──→ 标准二分（704）
  │
  ├─ 找 "第一个/最后一个 满足条件的位置" ──→ 左侧边界 或 右侧边界
  │
  ├─ 有序数组旋转了 ──→ 旋转数组二分（33/81/153/154）
  │     ├─ 无重复：nums[left] <= nums[mid]
  │     └─ 有重复：nums[left]==nums[mid]==nums[right] → left++
  │
  ├─ 二维矩阵（行首>上行尾）──→ 坐标映射一维二分（74）
  │
  └─ 求平方根/完全平方数 ──→ 特殊边界二分（69/367）
```

---

## 二、三种基础模板

### 模板1：标准二分（左闭右闭 [left, right]）

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1  // [left, right]
    for left <= right {             // 注意是 <=
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
**适用**：找存在且唯一的 target（704）

---

### 模板2：左侧边界（左闭右开 [left, right)）

```go
func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)      // [left, right)
    for left < right {               // 注意是 <
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid              // 保持 left 不动，收缩 right
        }
    }
    return left                      // left == right 时退出
}
```
**适用**：找第一个 >= target 的位置（35、34左侧）

---

### 模板3：右侧边界（左闭右开 [left, right)）

```go
func upperBound(nums []int, target int) int {
    left, right := 0, len(nums)      // [left, right)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1           // 找 > target，所以要越过 <= 的
        } else {
            right = mid
        }
    }
    return left                      // 第一个 > target 的位置
}
```
**适用**：找最后一个 <= target 的位置（34右侧）

---

## 三、旋转数组四姐妹

### 33/81：搜索旋转排序数组

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        // 左半边有序
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // 右半边有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

**81题多一步去重**：
```go
// 当无法判断哪半有序时，收缩左边界
if nums[left] == nums[mid] && nums[mid] == nums[right] {
    left++
}
```

### 153/154：找旋转数组最小值

```go
func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1     // 最小值在右半边
        } else if nums[mid] < nums[right] {
            right = mid        // 最小值在左半边（含mid）
        } else {              // nums[mid] == nums[right]，154多这一步
            right--            // 无法判断，收缩右边界
        }
    }
    return nums[left]
}
```

---

## 四、两个特殊变种

### 74：搜索二维矩阵

```go
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    left, right := 0, m*n-1
    for left <= right {
        mid := left + (right-left)/2
        row, col := mid/n, mid%n   // 关键映射
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

### 69/367：平方根系列

```go
func mySqrt(x int) int {
    if x < 2 {
        return x
    }
    left, right := 1, x/2
    for left <= right {
        mid := left + (right-left)/2
        if mid <= x/mid {           // 防止溢出，用 x/mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right                     // right 是最大的满足 mid<=x/mid 的
}
```

---

## 五、通用口诀

```
标准二分：left <= right，左闭右闭，target找到就返回
左边界：left < right，左闭右开，找 >= target
右边界：left < right，左闭右开，找 > target
旋转数组：先判断哪半有序，再判断 target 在不在有序的那半
防溢出：不用 mid*mid，用 x/mid 或 x/mid >= mid
```

---

## 六、34题完整答案（两次二分）

```go
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    // 找左侧边界
    left := lowerBound(nums, target)
    if left == len(nums) || nums[left] != target {
        return []int{-1, -1}
    }
    // 找右侧边界：target+1 的左侧边界 - 1
    right := lowerBound(nums, target+1) - 1
    return []int{left, right}
}

func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
```

---

## 七、一图总结

```
┌─────────────────────────────────────────────────────────┐
│                    二分查找速查卡                         │
├──────────────┬──────────────────┬──────────────────────┤
│ 标准二分     │ left <= right    │ 找存在的 target      │
│ 左边界       │ left < right, [) │ 第一个 >= target     │
│ 右边界       │ left < right, [) │ 第一个 > target      │
│ 旋转数组     │ 先判断哪半有序    │ 33/81/153/154       │
│ 矩阵映射     │ index → (row,col)│ 74 搜索二维矩阵      │
│ 防溢出       │ x/mid 替代 mid*mid│ 69/367 平方根      │
└──────────────┴──────────────────┴──────────────────────┘
```
