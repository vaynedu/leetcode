package leetcode

// 57.插入区间

import (
	"fmt"
	"testing"
)

/**
给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的
开始和结束，并且 intervals 按照 starti 升序排列。同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始
和结束。

 在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以
合并区间）。

 返回插入之后的 intervals。

 注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。



 示例 1：


输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]


 示例 2：


输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。




 提示：


 0 <= intervals.length <= 10⁴
 intervals[i].length == 2
 0 <= starti <= endi <= 10⁵
 intervals 根据 starti 按 升序 排列
 newInterval.length == 2
 0 <= start <= end <= 10⁵


 Related Topics 数组 👍 1016 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInsertInterval(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
