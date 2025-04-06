package leetcode

// 257.二叉树的所有路径

import (
	"fmt"
	"testing"
)

//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 1222 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}
	path += fmt.Sprintf("%d->", root.Val)
	if root.Right == nil && root.Left == nil {
		*paths = append(*paths, path[:len(path)-2])
		return
	}
	helper(root.Left, path, paths)
	helper(root.Right, path, paths)

	return
}
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	ans := make([]string, 0)
	helper(root, "", &ans)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreePaths(t *testing.T) {
	fmt.Println("come on")
	// 生成函数测试用例
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(binaryTreePaths(root))
}
