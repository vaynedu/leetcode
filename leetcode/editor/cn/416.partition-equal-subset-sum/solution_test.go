package leetcode

// 分割等和子集

import (
	"fmt"
	"testing"
)

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



 示例 1：


输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

 示例 2：


输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。




 提示：


 1 <= nums.length <= 200
 1 <= nums[i] <= 100


 Related Topics 数组 动态规划 👍 2293 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

func canPartition(nums []int) bool {
	// 计算出数组总和，然后遍历数组，如果某个数大于总和的一半，则直接返回false
	// 巧妙使用背包的思路
	// https://programmercarl.com/0416.%E5%88%86%E5%89%B2%E7%AD%89%E5%92%8C%E5%AD%90%E9%9B%86.html#_01%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- { // 每一个元素一定是不可重复放入，所以从大到小遍历
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

//leetcode submit region end(Prohibit modification and deletion)

//1.直接超时
//func backTrack(nums []int, index int, curSum int, sum int) bool {
//    if index >= len(nums) {
//        return false
//    }
//    if curSum*2 == sum {
//        return true
//    }
//    if curSum > sum/2 {
//        return false
//    }
//
//    // 选择当前 || // 不选择当前
//    return backTrack(nums, index+1, curSum+nums[index], sum) ||
//        backTrack(nums, index+1, curSum, sum)
//}
//func canPartition(nums []int) bool {
//    // 计算出数组总和，然后遍历数组，如果某个数大于总和的一半，则直接返回false
//    n := len(nums)
//    sum := 0
//    for i := 0; i < n; i++ {
//        sum += nums[i]
//    }
//    if sum%2 != 0 {
//        return false
//    }
//
//    return backTrack(nums, 0, 0, sum)
//}

// 2.使用dp数组
//
//	func canPartition(nums []int) bool {
//		// 计算出数组总和，然后遍历数组，如果某个数大于总和的一半，则直接返回false
//		n := len(nums)
//		sum := 0
//		for i := 0; i < n; i++ {
//			sum += nums[i]
//		}
//		if sum%2 != 0 {
//			return false
//		}
//
//		target := sum / 2
//		dp := make([]bool, target+1)
//		dp[0] = true
//		for i := 0; i < n; i++ {
//			for j := target; j >= nums[i]; j-- {
//				dp[j] = dp[j] || dp[j-nums[i]]
//			}
//		}
//		return dp[target]
//	}
func TestPartitionEqualSubsetSum(t *testing.T) {
	fmt.Println("come on baby !!!")
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
