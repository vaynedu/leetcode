package leetcode

// 有序三元组中的最大值 II

import (
	"fmt"
	"testing"
)

/**
给你一个下标从 0 开始的整数数组 nums 。

 请你从所有满足 i < j < k 的下标三元组 (i, j, k) 中，找出并返回下标三元组的最大值。如果所有满足条件的三元组的值都是负数，则返回 0 。


 下标三元组 (i, j, k) 的值等于 (nums[i] - nums[j]) * nums[k] 。



 示例 1：


输入：nums = [12,6,1,2,7]
输出：77
解释：下标三元组 (0, 2, 4) 的值是 (nums[0] - nums[2]) * nums[4] = 77 。
可以证明不存在值大于 77 的有序下标三元组。


 示例 2：


输入：nums = [1,10,3,4,19]
输出：133
解释：下标三元组 (1, 2, 4) 的值是 (nums[1] - nums[2]) * nums[4] = 133 。
可以证明不存在值大于 133 的有序下标三元组。


 示例 3：


输入：nums = [1,2,3]
输出：0
解释：唯一的下标三元组 (0, 1, 2) 的值是一个负数，(nums[0] - nums[1]) * nums[2] = -3 。因此，答案是 0 。




 提示：


 3 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁶


 Related Topics 数组 👍 43 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumTripletValue(nums []int) int64 {
	// 1.暴力解法,会超时
	//n := len(nums)
	//ans := int64(0)
	//for i := 0; i < n; i++ {
	//	for j := i + 1; j < n; j++ {
	//		for k := j + 1; k < n; k++ {
	//			if (nums[i]-nums[j])*nums[k] > 0 {
	//				ans = max(ans, int64((nums[i]-nums[j])*nums[k]))
	//			}
	//		}
	//	}
	//}
	//return ans

	//2.前缀数据组和
	n := len(nums)
	if n < 3 {
		return 0
	}

	//2.计算前缀最大值数组
	//maxLeft := make([]int, n)
	//maxLeft[0] = nums[0]
	//for i := 1; i < n; i++ {
	//	maxLeft[i] = max(maxLeft[i-1], nums[i])
	//}
	//
	//// 计算后缀最大值数组
	//maxRight := make([]int, n)
	//maxRight[n-1] = nums[n-1]
	//for i := n - 2; i >= 0; i-- {
	//	maxRight[i] = max(maxRight[i+1], nums[i])
	//}
	//
	//// 初始化答案
	//ans := int64(0)
	//
	//// 遍历每个可能的 j
	//for j := 1; j < n-1; j++ {
	//	// 计算 (maxLeft[j-1] - nums[j]) * maxRight[j+1]
	//	if val := int64((maxLeft[j-1] - nums[j]) * maxRight[j+1]); val > ans {
	//		ans = val
	//	}
	//}
	//return ans

	// 3. 前缀最大值和最大差值
	ans := 0
	maxValue := 0 // 前缀最大值
	maxDiff := 0  // 最大差值

	for _, num := range nums {
		ans = max(ans, maxDiff*num)
		maxValue = max(maxValue, num)
		maxDiff = max(maxDiff, maxValue-num)
	}

	return int64(ans)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumValueOfAnOrderedTripletIi(t *testing.T) {
	fmt.Println("come on baby !!!")
	nums := []int{12, 6, 1, 2, 7}
	ans := maximumTripletValue(nums)
	fmt.Println(ans)
}
