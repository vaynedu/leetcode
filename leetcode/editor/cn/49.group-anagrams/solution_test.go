package leetcode

// 49.字母异位词分组

import (
	"fmt"
	"sort"
	"testing"
)

/**
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。



 示例 1:


 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]


 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

 解释：


 在 strs 中没有字符串可以通过重新排列来形成 "bat"。
 字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
 字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。


 示例 2:


 输入: strs = [""]


 输出: [[""]]

 示例 3:


 输入: strs = ["a"]


 输出: [["a"]]



 提示：


 1 <= strs.length <= 10⁴
 0 <= strs[i].length <= 100
 strs[i] 仅包含小写字母


 Related Topics 数组 哈希表 字符串 排序 👍 2436 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	countMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		origin := strs[i]
		// 将字符串转为字节切片以便排序
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(a, b int) bool {
			return bytes[a] < bytes[b]
		})
		sortedStr := string(bytes)

		// 添加到对应键的列表中
		countMap[sortedStr] = append(countMap[sortedStr], origin)
	}

	// 构建结果数组
	res := make([][]string, 0, len(countMap))
	for _, vals := range countMap {
		res = append(res, vals)
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func TestGroupAnagrams(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
