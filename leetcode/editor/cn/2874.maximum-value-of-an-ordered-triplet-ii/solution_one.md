
### 示例

假设我们有一个数组 `nums = [12, 6, 1, 2, 7]`。

#### 1. 计算前缀最大值数组 (`maxLeft`)

前缀最大值数组 `maxLeft` 的每个元素 `maxLeft[i]` 表示从 `nums[0]` 到 `nums[i]` 的最大值。

| i  | nums[i] | maxLeft[i] |
|----|---------|------------|
| 0  | 12      | 12         |
| 1  | 6       | 12         |
| 2  | 1       | 12         |
| 3  | 2       | 12         |
| 4  | 7       | 12         |

计算过程：
- `maxLeft[0] = nums[0] = 12`
- `maxLeft[1] = max(maxLeft[0], nums[1]) = max(12, 6) = 12`
- `maxLeft[2] = max(maxLeft[1], nums[2]) = max(12, 1) = 12`
- `maxLeft[3] = max(maxLeft[2], nums[3]) = max(12, 2) = 12`
- `maxLeft[4] = max(maxLeft[3], nums[4]) = max(12, 7) = 12`

#### 2. 计算后缀最大值数组 (`maxRight`)

后缀最大值数组 `maxRight` 的每个元素 `maxRight[i]` 表示从 `nums[i]` 到 `nums[n-1]` 的最大值。

| i  | nums[i] | maxRight[i] |
|----|---------|-------------|
| 0  | 12      | 12          |
| 1  | 6       | 12          |
| 2  | 1       | 12          |
| 3  | 2       | 7           |
| 4  | 7       | 7           |

计算过程：
- `maxRight[4] = nums[4] = 7`
- `maxRight[3] = max(maxRight[4], nums[3]) = max(7, 2) = 7`
- `maxRight[2] = max(maxRight[3], nums[2]) = max(7, 1) = 7`
- `maxRight[1] = max(maxRight[2], nums[1]) = max(7, 6) = 7`
- `maxRight[0] = max(maxRight[1], nums[0]) = max(7, 12) = 12`

#### 3. 计算最大值

现在我们使用 `maxLeft` 和 `maxRight` 来计算最大值 `(maxLeft[j-1] - nums[j]) * maxRight[j+1]`。

| j  | maxLeft[j-1] | nums[j] | maxRight[j+1] | (maxLeft[j-1] - nums[j]) * maxRight[j+1] |
|----|--------------|---------|---------------|------------------------------------------|
| 1  | 12           | 6       | 7             | (12 - 6) * 7 = 42                        |
| 2  | 12           | 1       | 7             | (12 - 1) * 7 = 77                        |
| 3  | 12           | 2       | 7             | (12 - 2) * 7 = 70                        |

计算过程：
- 对于 `j = 1`：
    - `maxLeft[0] = 12`
    - `nums[1] = 6`
    - `maxRight[2] = 7`
    - `(maxLeft[0] - nums[1]) * maxRight[2] = (12 - 6) * 7 = 42`
- 对于 `j = 2`：
    - `maxLeft[1] = 12`
    - `nums[2] = 1`
    - `maxRight[3] = 7`
    - `(maxLeft[1] - nums[2]) * maxRight[3] = (12 - 1) * 7 = 77`
- 对于 `j = 3`：
    - `maxLeft[2] = 12`
    - `nums[3] = 2`
    - `maxRight[4] = 7`
    - `(maxLeft[2] - nums[3]) * maxRight[4] = (12 - 2) * 7 = 70`

最终的最大值是 `77`。

### 图示

以下是 `maxLeft` 和 `maxRight` 数组的图示：

#### 前缀最大值数组 (`maxLeft`)

```
nums:     [12, 6, 1, 2, 7]
maxLeft:  [12, 12, 12, 12, 12]
```


#### 后缀最大值数组 (`maxRight`)

```
nums:     [12, 6, 1, 2, 7]
maxRight: [12, 7, 7, 7, 7]
```


### 完整代码

以下是完整的代码，包含详细的注释和计算过程：

```go
package leetcode

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
	n := len(nums)
	if n < 3 {
		return 0
	}

	// 计算前缀最大值数组
	maxLeft := make([]int, n)
	maxLeft[0] = nums[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], nums[i])
	}

	// 计算后缀最大值数组
	maxRight := make([]int, n)
	maxRight[n-1] = nums[n-1]
	for i := n-2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], nums[i])
	}

	// 初始化答案
	ans := int64(0)

	// 遍历每个可能的 j
	for j := 1; j < n-1; j++ {
		// 计算 (maxLeft[j-1] - nums[j]) * maxRight[j+1]
		if val := int64((maxLeft[j-1] - nums[j]) * maxRight[j+1]); val > ans {
			ans = val
		}
	}

	return ans
}

// 辅助函数：返回两个整数中的最大值
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
	fmt.Println(ans) // 输出: 77
}
```