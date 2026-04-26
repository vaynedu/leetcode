package interview

/*
617. 合并二叉树

给你两棵二叉树 root1 和 root2
如果节点重叠，则将两个节点的值相加；否则返回该节点
可以原地合并（返回 root1）或新建一棵树

时间：O(min(m,n)) | 空间：O(min(h1,h2))
*/

// MergeTrees 递归合并
func MergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 两个节点都存在，合并值
	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)
	return root1
}
