package leetcode

// 76.最小覆盖子串

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



 注意：


 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 如果 s 中存在这样的子串，我们保证它是唯一的答案。




 示例 1：


输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。


 示例 2：


输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。


 示例 3:


输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。



 提示：


 m == s.length
 n == t.length
 1 <= m, n <= 10⁵
 s 和 t 由英文字母组成



进阶：你能设计一个在
o(m+n) 时间内解决此问题的算法吗？

 👍 3375 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// need 记录字符串t中每个字符需要的数量
	need := make(map[byte]int)
	// window 记录当前窗口中每个字符出现的次数
	window := make(map[byte]int)

	// 初始化need map
	for i := range t {
		need[t[i]]++
	}

	// valid 记录窗口中满足need条件的字符种类数
	valid := 0

	// 记录最小覆盖子串的起始索引及长度
	start := 0
	minLen := len(s) + 1 // 初始化为比s长度大1的值

	// 滑动窗口的左右指针
	left, right := 0, 0

	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 扩大窗口
		right++

		// 进行窗口内数据的一系列更新
		if need[c] > 0 {
			window[c]++
			// 如果窗口中字符c的数量达到需要的数量，则valid加1
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++

			// 进行窗口内数据的一系列更新
			if need[d] > 0 {
				// 如果窗口中字符d的数量刚好等于需要的数量，移除后就不满足了，valid减1
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	// 返回最小覆盖子串
	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumWindowSubstring(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
