package leetcode

// 376.摆动序列

import (
	"fmt"
	"testing"
)

/**
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作
摆动序列。


 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差
值为零。


 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。

 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。



 示例 1：


输入：nums = [1,7,4,9,2,5]
输出：6
解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。


 示例 2：


输入：nums = [1,17,5,10,13,15,10,5,16,8]
输出：7
解释：这个序列包含几个长度为 7 摆动序列。
其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。


 示例 3：


输入：nums = [1,2,3,4,5,6,7,8,9]
输出：2




 提示：


 1 <= nums.length <= 1000
 0 <= nums[i] <= 1000




 进阶：你能否用 O(n) 时间复杂度完成此题?

 Related Topics 贪心 数组 动态规划 👍 1223 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	n := len(nums)
	ans := 1 // 记录峰值个数，序列默认序列最右边有一个峰值
	preDiff := 0
	curDiff := 0 // 当前一对差值
	for i := 0; i < n-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (preDiff >= 0 && curDiff < 0) ||
			(preDiff <= 0 && curDiff > 0) {
			ans++
			preDiff = curDiff
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestWiggleSubsequence(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
