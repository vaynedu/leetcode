package interview

import "sort"

// 15. 三数之和
// 双指针：排序后，对每个数用双指针找两数之和为零
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    n := len(nums)
    for i := 0; i < n-2 && nums[i] <= 0; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left, right := i+1, n-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return result
}
