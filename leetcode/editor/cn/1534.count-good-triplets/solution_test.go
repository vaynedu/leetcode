package leetcode

// 1534.统计好三元组

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 arr ，以及 a、b 、c 三个整数。请你统计其中好三元组的数量。

 如果三元组 (arr[i], arr[j], arr[k]) 满足下列全部条件，则认为它是一个 好三元组 。


 0 <= i < j < k < arr.length
 |arr[i] - arr[j]| <= a
 |arr[j] - arr[k]| <= b
 |arr[i] - arr[k]| <= c


 其中 |x| 表示 x 的绝对值。

 返回 好三元组的数量 。



 示例 1：

 输入：arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
输出：4
解释：一共有 4 个好三元组：[(3,0,1), (3,0,1), (3,1,1), (0,1,1)] 。


 示例 2：

 输入：arr = [1,1,2,2,3], a = 0, b = 0, c = 1
输出：0
解释：不存在满足所有条件的三元组。




 提示：


 3 <= arr.length <= 100
 0 <= arr[i] <= 1000
 0 <= a, b, c <= 1000


 Related Topics 数组 枚举 👍 145 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountGoodTriplets(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
