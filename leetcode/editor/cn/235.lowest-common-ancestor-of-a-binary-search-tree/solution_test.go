package leetcode

// 235.二叉搜索树的最近公共祖先

import (
	"testing"
)

/**
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个
节点也可以是它自己的祖先）。”

 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]





 示例 1:

 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。


 示例 2:

 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。



 说明:


 所有节点的值都是唯一的。
 p、q 为不同节点且均存在于给定的二叉搜索树中。


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1314 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val { // p 和 q 都在左子树，向左子树搜索
		root = lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val { // p 和 q 都在右子树，向右子树搜索
		root = lowestCommonAncestor(root.Right, p, q)
	}

	// p 和 q 分别在左右子树，或者其中一个节点就是当前根节点，当前根节点就是最近公共祖先
	return root

}

//leetcode submit region end(Prohibit modification and deletion)

//利用一个循环遍历二叉搜索树。
//如果 p 和 q 的值都小于当前根节点的值，说明它们都在左子树，将根节点更新为左子节点。
//如果 p 和 q 的值都大于当前根节点的值，说明它们都在右子树，将根节点更新为右子节点。
//否则，说明 p 和 q 分别在左右子树，或者其中一个节点就是当前根节点，当前根节点就是最近公共祖先，返回该节点。

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {
	// 构建示例二叉搜索树
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
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
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 8},
			want: &TreeNode{Val: 6},
		},
		{
			name: "Example 2",
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 为了比较方便，这里先获取实际结果和预期结果的值
			got := lowestCommonAncestor(root, findNode(root, tt.p.Val), findNode(root, tt.q.Val))
			if got == nil || got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}

// findNode 在二叉搜索树中查找值为 val 的节点
func findNode(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
