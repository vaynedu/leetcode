package leetcode

// 71.简化路径

import (
	"strings"
	"testing"
)

/**
给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为 更加简洁的规范路径。

 在 Unix 风格的文件系统中规则如下：


 一个点 '.' 表示当前目录本身。
 此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
 任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
 任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。


 返回的 简化路径 必须遵循下述格式：


 始终以斜杠 '/' 开头。
 两个目录名之间必须只有一个斜杠 '/' 。
 最后一个目录名（如果存在）不能 以 '/' 结尾。
 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。


 返回简化后得到的 规范路径 。



 示例 1：


 输入：path = "/home/"


 输出："/home"

 解释：

 应删除尾随斜杠。

 示例 2：


 输入：path = "/home//foo/"


 输出："/home/foo"

 解释：

 多个连续的斜杠被单个斜杠替换。

 示例 3：


 输入：path = "/home/user/Documents/../Pictures"


 输出："/home/user/Pictures"

 解释：

 两个点 ".." 表示上一级目录（父目录）。

 示例 4：


 输入：path = "/../"


 输出："/"

 解释：

 不可能从根目录上升一级目录。

 示例 5：


 输入：path = "/.../a/../b/c/../d/./"


 输出："/.../b/d"

 解释：

 "..." 在这个问题中是一个合法的目录名。



 提示：


 1 <= path.length <= 3000
 path 由英文字母，数字，'.'，'/' 或 '_' 组成。
 path 是一个有效的 Unix 风格绝对路径。


 Related Topics 栈 字符串 👍 843 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	stack := []string{}

	components := strings.Split(path, "/")

	for _, component := range components {
		// 空字符串或 '.' 表示当前目录，跳过处理
		if component == "" || component == "." {
			continue
		}
		// '..' 表示返回上级目录
		if component == ".." {
			// 如果栈不为空，则弹出栈顶元素
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, component)
	}

	return "/" + strings.Join(stack, "/")

}

//leetcode submit region end(Prohibit modification and deletion)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Example 1",
			path:     "/home/",
			expected: "/home",
		},
		{
			name:     "Example 2",
			path:     "/home//foo/",
			expected: "/home/foo",
		},
		{
			name:     "Example 3",
			path:     "/home/user/Documents/../Pictures",
			expected: "/home/user/Pictures",
		},
		{
			name:     "Example 4",
			path:     "/../",
			expected: "/",
		},
		{
			name:     "Example 5",
			path:     "/.../a/../b/c/../d/./",
			expected: "/.../b/d",
		},
		{
			name:     "Root path",
			path:     "/",
			expected: "/",
		},
		{
			name:     "Multiple dots",
			path:     "/a/./b/../../c/",
			expected: "/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifyPath(tt.path)
			if result != tt.expected {
				t.Errorf("simplifyPath(%q) = %q; expected %q", tt.path, result, tt.expected)
			}
		})
	}
}
