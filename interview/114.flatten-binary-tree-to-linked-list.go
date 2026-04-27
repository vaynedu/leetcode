package interview

// 114. 二叉树展开为链表
// 递归：将左子树展开接到根右，左子树置空，右子树接到最右节点
func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    flatten(root.Left)
    flatten(root.Right)
    if root.Left == nil {
        return
    }
    // 保存右子树
    right := root.Right
    // 左子树展开放到右边
    root.Right = root.Left
    root.Left = nil
    // 找到最右节点，连接原右子树
    cur := root.Right
    for cur.Right != nil {
        cur = cur.Right
    }
    cur.Right = right
}
