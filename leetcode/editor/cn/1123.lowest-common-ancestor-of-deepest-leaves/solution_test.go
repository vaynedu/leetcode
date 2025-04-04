package leetcode

// 最深叶节点的最近公共祖先

import (
	"fmt"
	"testing"
)

/**
给你一个有根节点
 root 的二叉树，返回它 最深的叶节点的最近公共祖先 。

 回想一下：


 叶节点 是二叉树中没有子节点的节点
 树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
 如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。




 示例 1：


输入：root = [3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：我们返回值为 2 的节点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的节点。
注意，节点 6、0 和 8 也是叶节点，但是它们的深度是 2 ，而节点 7 和 4 的深度是 3 。


 示例 2：


输入：root = [1]
输出：[1]
解释：根节点是树中最深的节点，它是它本身的最近公共祖先。


 示例 3：


输入：root = [0,1,3,null,2]
输出：[2]
解释：树中最深的叶节点是 2 ，最近公共祖先是它自己。



 提示：


 树中的节点数将在
 [1, 1000] 的范围内。
 0 <= Node.val <= 1000
 每个节点的值都是 独一无二 的。




 注意：本题与力扣 865 重复：https://leetcode-cn.com/problems/smallest-subtree-with-all-the-
deepest-nodes/

 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 319 👎 0

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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 题目思路
	// 获取左右子树的深度
	// left == right ，返回当前节点和左右子树的节点
	// left < right ，返回右子树的节点, 继续向下找
	// left > right ，返回左子树的节点, 继续向下找

	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left == right {
		return root
	} else if left > right {
		return lcaDeepestLeaves(root.Left)
	} else {
		return lcaDeepestLeaves(root.Right)
	}
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestLowestCommonAncestorOfDeepestLeaves(t *testing.T) {
	// 测试用例1
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 0}
	root1.Right.Right = &TreeNode{Val: 8}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	if result := lcaDeepestLeaves(root1); result.Val != 2 {
		t.Errorf("Test case 1 failed, expected 2, got %d", result.Val)
	}

	// 测试用例2
	root2 := &TreeNode{Val: 1}
	if result := lcaDeepestLeaves(root2); result.Val != 1 {
		t.Errorf("Test case 2 failed, expected 1, got %d", result.Val)
	}

	// 测试用例3
	root3 := &TreeNode{Val: 0}
	root3.Left = &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 3}
	root3.Left.Right = &TreeNode{Val: 2}
	if result := lcaDeepestLeaves(root3); result.Val != 2 {
		t.Errorf("Test case 3 failed, expected 2, got %d", result.Val)
	}

	// 测试用例4: 单个叶子节点
	root4 := &TreeNode{Val: 0}
	root4.Left = &TreeNode{Val: 1}
	root4.Left.Left = &TreeNode{Val: 2}
	if result := lcaDeepestLeaves(root4); result.Val != 2 {
		t.Errorf("Test case 4 failed, expected 2, got %d", result.Val)
	}

	fmt.Println("All test cases passed!")
}
