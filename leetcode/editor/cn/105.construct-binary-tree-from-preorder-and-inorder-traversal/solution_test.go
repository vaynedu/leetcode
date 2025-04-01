package leetcode

// 从前序与中序遍历序列构造二叉树

import (
	"fmt"
	"testing"
)

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。
//
//
//
// 示例 1:
//
//
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//
//
//
//
// 提示:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder 和 inorder 均 无重复 元素
// inorder 均出现在 preorder
// preorder 保证 为二叉树的前序遍历序列
// inorder 保证 为二叉树的中序遍历序列
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 2507 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	inorderIndexMap := make(map[int]int)
	for i, v := range inorder {
		inorderIndexMap[v] = i
	}

	var buildTreeHelper func(preStart, preEnd, inStart, inEnd int) *TreeNode
	buildTreeHelper = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		rootIndex := inorderIndexMap[rootVal]
		leftLen := rootIndex - inStart
		root.Left = buildTreeHelper(preStart+1, preStart+leftLen, inStart, rootIndex-1)
		root.Right = buildTreeHelper(preStart+leftLen+1, preEnd, rootIndex+1, inEnd)
		return root
	}

	return buildTreeHelper(0, len(preorder)-1, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
	// 测试用例 1
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	result1 := buildTree(preorder1, inorder1)
	t.Logf("Test Case 1 Result: %+v", result1)

	// 测试用例 2
	preorder2 := []int{-1}
	inorder2 := []int{-1}
	result2 := buildTree(preorder2, inorder2)
	t.Logf("Test Case 2 Result: %+v", result2)

	// 测试用例 3：空数组
	preorder3 := []int{}
	inorder3 := []int{}
	result3 := buildTree(preorder3, inorder3)
	t.Logf("Test Case 3 Result: %+v", result3)

	// 测试用例 4：单节点树
	preorder4 := []int{1}
	inorder4 := []int{1}
	result4 := buildTree(preorder4, inorder4)
	t.Logf("Test Case 4 Result: %+v", result4)

	fmt.Println("All test cases passed!")
}
