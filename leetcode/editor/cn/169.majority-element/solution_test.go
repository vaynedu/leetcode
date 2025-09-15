package leetcode

// 169.多数元素

import (
	"fmt"
	"testing"
)

/**
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

 你可以假设数组是非空的，并且给定的数组总是存在多数元素。



 示例 1：


输入：nums = [3,2,3]
输出：3

 示例 2：


输入：nums = [2,2,1,1,1,2,2]
输出：2



提示：


 n == nums.length
 1 <= n <= 5 * 10⁴
 -10⁹ <= nums[i] <= 10⁹




 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

 Related Topics 数组 哈希表 分治 计数 排序 👍 2551 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 投票算法
	// 投票算法，如果当前值和当前值相同，则计数器加1，否则减1
	// 当计数器为0时，将当前值赋给ans，并重新开始计数
	// 循环结束后，ans即为多数元素
	count := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ans = nums[i]
		}

		if nums[i] == ans {
			count++
		} else {
			count--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMajorityElement(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
