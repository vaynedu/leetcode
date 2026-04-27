package interview

// 173. 二叉搜索树迭代器
// 中序遍历：用栈模拟，next/hasNext O(1) 均摊
type BSTIterator struct {
    stack []*TreeNode
    node  *TreeNode
}

func NewBSTIterator(root *TreeNode) BSTIterator {
    return BSTIterator{node: root}
}

func (this *BSTIterator) Next() int {
    for this.node != nil {
        this.stack = append(this.stack, this.node)
        this.node = this.node.Left
    }
    this.node = this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    val := this.node.Val
    this.node = this.node.Right
    return val
}

func (this *BSTIterator) HasNext() bool {
    return this.node != nil || len(this.stack) > 0
}
