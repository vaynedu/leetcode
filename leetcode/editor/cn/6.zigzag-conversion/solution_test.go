package leetcode

// 6.Z 字形变换

import (
	"fmt"
	"testing"
)

/**
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：


P   A   H   N
A P L S I I G
Y   I   R

 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

 请你实现这个将字符串进行指定行数变换的函数：


string convert(string s, int numRows);



 示例 1：


输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"


示例 2：


输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I


 示例 3：


输入：s = "A", numRows = 1
输出："A"




 提示：


 1 <= s.length <= 1000
 s 由英文字母（小写和大写）、',' 和 '.' 组成
 1 <= numRows <= 1000


 Related Topics 字符串 👍 2576 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {

	// 特殊情况：行数为1时直接返回原字符串
	if numRows == 1 {
		return s
	}

	// 创建行数组存储每行的字符
	rows := make([][]byte, numRows)
	for i := range rows {
		rows[i] = make([]byte, 0)
	}

	// 当前行号和方向（向下为1，向上为-1）
	currentRow := 0
	direction := 1

	// 遍历字符串中的每个字符
	for i := range s {
		// 将字符添加到对应行
		rows[currentRow] = append(rows[currentRow], s[i])

		// 更新下一行的方向和行号
		if currentRow == 0 {
			direction = 1 // 到达顶部，开始向下
		} else if currentRow == numRows-1 {
			direction = -1 // 到达底部，开始向上
		}

		currentRow += direction
	}

	// 拼接所有行的结果
	result := ""
	for i := range rows {
		result += string(rows[i])
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZigzagConversion(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
