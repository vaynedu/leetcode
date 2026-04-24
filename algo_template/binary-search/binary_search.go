package binary_search

// ============================================================
// 二分查找模板 - 三种写法
// ============================================================

// --------------------------------------------------
// 模板1: 标准二分查找（左闭右闭 [left, right]）
// 适用于：查找存在且唯一的元素
// --------------------------------------------------

// BinarySearch 704. 二分查找
// 标准模板：while left <= right，左闭右闭
func BinarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
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

// --------------------------------------------------
// 模板2: 寻找左侧边界的二分查找（左闭右开 [left, right)）
// 适用于：查找第一个 >= target 的位置
// --------------------------------------------------

// LowerBound 35. 搜索插入位置（也是左边界）
// 返回第一个 >= target 的下标
func LowerBound(nums []int, target int) int {
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

// FirstBadVersion 278. 第一个错误的版本
// 版本正确返回 false，第一个错误版本返回 true
func FirstBadVersion(n int) int {
    left, right := 1, n
    for left < right {
        mid := left + (right-left)/2
        if IsBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

// --------------------------------------------------
// 模板3: 寻找右侧边界的二分查找（左闭右开 [left, right)）
// 适用于：查找最后一个 <= target 的位置
// --------------------------------------------------

// UpperBound 查找第一个 > target 的位置（右边界）
func UpperBound(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

// --------------------------------------------------
// 类型4: 求平方根（特殊边界）
// --------------------------------------------------

// MySqrt 69. x 的平方根（向下取整）
func MySqrt(x int) int {
    if x < 2 {
        return x
    }
    left, right := 1, x/2
    for left <= right {
        mid := left + (right-left)/2
        // 防止 mid*mid 溢出，用 x/mid
        if mid <= x/mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right
}

// ValidPerfectSquare 367. 有效的完全平方数
func ValidPerfectSquare(num int) bool {
    if num < 2 {
        return true
    }
    left, right := 1, num/2
    for left <= right {
        mid := left + (right-left)/2
        if mid <= num/mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left*left == num || right*right == num
}

// --------------------------------------------------
// 类型5: 搜索旋转排序数组
// --------------------------------------------------

// SearchRotated 33. 搜索旋转排序数组
// 时间 O(log n)，空间 O(1)
func SearchRotated(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
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

// FindMin 153. 寻找旋转排序数组中的最小值
func FindMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}

// FindPeakElement 162. 寻找峰值
// 峰值是比相邻元素都大的元素，数组两端可以看作负无穷
func FindPeakElement(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < nums[mid+1] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

// --------------------------------------------------
// 辅助函数（用于测试）
// --------------------------------------------------

// IsBadVersion 是 LeetCode 官方题提供的函数
// 这里提供模拟实现
func IsBadVersion(version int) bool {
    // 模拟：假设版本 5 是第一个错误版本
    return version >= 5
}
