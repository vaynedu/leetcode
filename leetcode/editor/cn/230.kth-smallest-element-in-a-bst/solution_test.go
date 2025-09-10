package leetcode

// 230.二叉搜索树中第 K 小的元素

import (
	"fmt"
	"testing"

	. "leetcode/model"
)

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。



 示例 1：


输入：root = [3,1,4,null,2], k = 1
输出：1


 示例 2：


输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3






 提示：


 树中的节点数为 n 。
 1 <= k <= n <= 10⁴
 0 <= Node.val <= 10⁴




 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1022 👎 0

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
func kthSmallest(root *TreeNode, k int) int {
	// 方法一：递归中序遍历
	count := 0
	result := 0

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}

		// 遍历左子树
		inorder(node.Left)

		// 访问当前节点
		count++
		if count == k {
			result = node.Val
			return
		}

		// 遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return result

	// 方法二：迭代中序遍历（注释掉的方法）
	/*
	   stack := []*TreeNode{}
	   current := root
	   count := 0

	   for current != nil || len(stack) > 0 {
	       // 一直向左走到底
	       for current != nil {
	           stack = append(stack, current)
	           current = current.Left
	       }

	       // 弹出栈顶元素
	       current = stack[len(stack)-1]
	       stack = stack[:len(stack)-1]

	       count++
	       if count == k {
	           return current.Val
	       }

	       // 转向右子树
	       current = current.Right
	   }

	   return 0
	*/
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthSmallestElementInABst(t *testing.T) {
	fmt.Println("Testing kthSmallest...")

	// 构建测试用例1: [3,1,4,null,2], k=1
	//       3
	//      / \
	//     1   4
	//      \
	//       2
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 2}

	tests := []struct {
		root     *TreeNode
		k        int
		expected int
		name     string
	}{
		{root1, 1, 1, "Test case 1"},
		{root1, 2, 2, "Test case 2"},
		{root1, 3, 3, "Test case 3"},
		{root1, 4, 4, "Test case 4"},
	}

	for _, test := range tests {
		result := kthSmallest(test.root, test.k)
		if result != test.expected {
			t.Errorf("%s failed: expected %d, got %d", test.name, test.expected, result)
		} else {
			fmt.Printf("%s passed: got %d\n", test.name, result)
		}
	}

	// 构建测试用例2: [5,3,6,2,4,null,null,1], k=3
	//         5
	//       /   \
	//      3     6
	//     / \
	//    2   4
	//   /
	//  1
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 6}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 1}

	result := kthSmallest(root2, 3)
	expected := 3
	if result != expected {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected, result)
	} else {
		fmt.Printf("Test case 2 passed: got %d\n", result)
	}
}
