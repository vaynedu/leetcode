package leetcode

// 189.轮转数组

import (
	"fmt"
	"testing"
)

/**
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



 示例 1:


输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]


 示例 2:


输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]



 提示：


 1 <= nums.length <= 10⁵
 -2³¹ <= nums[i] <= 2³¹ - 1
 0 <= k <= 10⁵




 进阶：


 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


 Related Topics 数组 数学 双指针 👍 2435 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func rotate(nums []int, k int) {
	// 三次反转法的数据原理基于置换分解思想
	//这是三次反转法（也称为"环形替换法"的变种），核心思想是通过三次反转操作实现数组轮转：
	//第一步：对整个数组进行反转
	//第二步：对前k个元素进行反转
	//第三步：对后n-k个元素进行反
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRotateArray(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
