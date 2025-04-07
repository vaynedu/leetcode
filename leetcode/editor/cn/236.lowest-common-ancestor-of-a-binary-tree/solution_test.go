package leetcode

import "testing"

// 236.二叉树的最近公共祖先

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个
节点也可以是它自己的祖先）。”



 示例 1：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。


 示例 2：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。


 示例 3：


输入：root = [1,2], p = 1, q = 2
输出：1




 提示：


 树中节点数目在范围 [2, 10⁵] 内。
 -10⁹ <= Node.val <= 10⁹
 所有 Node.val 互不相同 。
 p != q
 p 和 q 均存在于给定的二叉树中。


 Related Topics 树 深度优先搜索 二叉树 👍 2954 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点为空，或者当前节点就是 p 或者 q，直接返回当前节点
	if root == nil || root == p || root == q {
		return root
	}
	// 递归调用左子树
	left := lowestCommonAncestor(root.Left, p, q)
	// 递归调用右子树
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果 left 和 right 都不为空，说明 p 和 q 分别位于当前节点的左右子树中，当前节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	// 如果 left 不为空，right 为空，说明 p 和 q 都在左子树中，left 就是最近公共祖先
	if left != nil {
		return left
	}
	// 否则，right 就是最近公共祖先
	return right
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// TestLowestCommonAncestor 测试 lowestCommonAncestor 函数
func TestLowestCommonAncestor(t *testing.T) {
	// 构建示例二叉树
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			name: "Example 1",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 1},
			want: &TreeNode{Val: 3},
		},
		{
			name: "Example 2",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pNode := findNode(root, tt.p.Val)
			qNode := findNode(root, tt.q.Val)
			wantNode := findNode(root, tt.want.Val)
			got := lowestCommonAncestor(root, pNode, qNode)
			if got != wantNode {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, wantNode.Val)
			}
		})
	}
}

// findNode 在二叉树中查找值为 val 的节点
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	return findNode(root.Right, val)
}
